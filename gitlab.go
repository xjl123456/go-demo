package main

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
	"os/exec"
)

type Client interface {
	Clone(filepath string)
	Push(gitAuth *http.BasicAuth)
	OpenWorkspace(filepath string)
	Commit(msg string)
	Add()
	Checkout()
}

type client struct {
	Url          string          //git的repo地址
	AuthToken    string          // git的accessTocken
	AuthUser     string          //用户名
	AuthPassword string          //用户密码
	BranchName   string          //分支名称
	TagName      string          //tag名称
	RemoteName   string          //远程库的名称
	Repository   *git.Repository //操作时的repo对象
	Worktree     *git.Worktree   //操作时的worktree对象
	Err          error           //收集的错误信息
}

func NewGitApiClient(url, token, user, pwd, brancchName, tagName, remoteName string) Client {
	return &client{
		Url:          url,
		AuthToken:    token,
		AuthUser:     user,
		AuthPassword: pwd,
		BranchName:   brancchName,
		TagName:      tagName,
		RemoteName:   remoteName,
	}
}

func main1() {
	//url := "http://172.20.2.149/gitlab-test/forkpath-unicom123-region-MFkmGj-test-all-aaa-22-latest"
	url := "http://172.20.2.149/faas-samples/aaa.git"
	authToken := "cDsXPQDsx-_eszxZeya7"
	authUser := "xijl7"
	authPassword := ""
	tagName := "1"
	branchName := "master"
	remoteName := "origin"
	//var gitAuth = &http.BasicAuth{Username:"xijl7", Password: "cDsXPQDsx-_eszxZeya7"}

	gitClinet := NewGitApiClient(url, authToken, authUser, authPassword, branchName, tagName, remoteName)
	gitClinet.OpenWorkspace("D:\\desktop\\22222")
	gitClinet.Checkout()
	//gitClinet.Add()
	//gitClinet.Commit("test")
	//gitClinet.Push(gitAuth)

}

func (g *client) Clone(filepath string) {
	var referenceName plumbing.ReferenceName
	//判断是分支还是tag
	if g.TagName == "" {
		referenceName = plumbing.NewBranchReferenceName(g.BranchName)
	} else {
		referenceName = plumbing.NewTagReferenceName(g.TagName)
	}
	//执行clone操作
	g.Repository, g.Err = git.PlainClone(filepath, false, &git.CloneOptions{
		URL: g.Url,
		Auth: &http.BasicAuth{
			Username: g.AuthUser,
			Password: g.AuthToken,
		},
		ReferenceName: referenceName,
		Progress:      os.Stdout,
	})
	//clone的代码权限为1000.1000(webIde的pod用户是1000)
	cmd := exec.Command("chown", "-R", "1000.1000", filepath)
	cmd.Start()
}

func (g *client) Push(gitAuth *http.BasicAuth) {
	//验证Repository是否为空
	if g.Repository == nil {
		g.Err = errors.New("init Repository first")
		return

	}

	g.Err = g.Repository.Push(&git.PushOptions{
		RemoteName:        g.RemoteName,
		RefSpecs:          nil,
		Auth:              gitAuth,
		Progress:          os.Stdout,
		Prune:             false,
		Force:             false,
		InsecureSkipTLS:   false,
		CABundle:          nil,
		RequireRemoteRefs: nil,
	})
	fmt.Println(g.Err)
}

//OpenWorkspace 打开已有的工作空间，已有.git文件
//filepath 是文件系统的地址
func (g *client) OpenWorkspace(filepath string) {
	//获取到git的repo对象，初始化到结构体
	g.Repository, g.Err = git.PlainOpen(filepath)
	//获取到gitworktree对象，初始化到结构体
	g.Worktree, g.Err = g.Repository.Worktree()
}

/*Commit git commit的操作
参数 msg 提交的信息 -m 里的内容
需要 Worktree已经初始化
*/
func (g *client) Commit(msg string) {
	//验证worktree是否为空
	if g.Worktree == nil {
		g.Err = errors.New("init worktree first")
		return

	}
	//commit 操作
	_, g.Err = g.Worktree.Commit(msg, &git.CommitOptions{})
	checkIfError(g.Err)
	return
}

/*Add git add的操作
需要 Worktree已经初始化
*/
func (g *client) Add() {
	//验证worktree是否为空
	if g.Worktree == nil {
		g.Err = errors.New("init worktree first")
		return
	}
	// exec git add -A
	g.Err = g.Worktree.AddWithOptions(&git.AddOptions{
		All: true,
	})
	// check delete file to add
	status_, _ := g.Worktree.Status()
	if status_.IsClean() == false {
		for key := range status_ {
			g.Worktree.Add(key)
		}
	}

	checkIfError(g.Err)
	return
}

func checkIfError(err error) {
	if err == nil {
		return
	} else {
		//todo 错误处理，目前就是打印
		fmt.Println(err)
	}

}

func (g *client) Checkout() {
	//验证worktree是否为空
	if g.Worktree == nil {
		g.Err = errors.New("init worktree first")
		return
	}
	var referenceName plumbing.ReferenceName

	if g.TagName == "" {
		referenceName = plumbing.NewBranchReferenceName(g.BranchName)
	} else {
		referenceName = plumbing.NewTagReferenceName(g.TagName)
	}
	g.Err = g.Worktree.Checkout(&git.CheckoutOptions{Branch: referenceName, Create: true})
	checkIfError(g.Err)
	return
}

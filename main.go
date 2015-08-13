package main

import (
  git "github.com/libgit2/git2go"
  "log"
  "C"
)

func credentialsCallback(url string, username string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {
  ret, cred := git.NewCredSshKey("git", "/Users/realbot/.ssh/id_rsa.pub", "/Users/realbot/.ssh/id_rsa", "")
  return git.ErrorCode(ret), &cred
}

// Made this one just return 0 during troubleshooting...
func certificateCheckCallback(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
    return 0
}

func main() {

  cbs := &git.RemoteCallbacks{
      CredentialsCallback:      credentialsCallback,
      CertificateCheckCallback: certificateCheckCallback,
  }

  cloneOptions := &git.CloneOptions{}
  //cloneOptions.RemoteCallbacks = cbs
  _ = cbs

  repo, err := git.Clone("/Users/realbot/Dropbox/gitos/veles.git", "veles", cloneOptions)
  if err != nil {
    log.Panic(err)
  }

  log.Print(repo)
}
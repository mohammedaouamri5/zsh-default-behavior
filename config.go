package main





type Start struct {}
type End struct {}


func (Start) IsGithubRipo(commend string) {
	Run("git clone " , commend)
} 


func (Start) IsImg(commend string)  {
	Run("chafa " , commend)
}





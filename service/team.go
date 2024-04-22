package service

import (
	"SCIProj/dao"
	"SCIProj/model"
)

func GetStudentNumsByCid(cid string) (num int, err error) {
	num, err = dao.GetStudentNumsByCid(cid)
	if err != nil {
		return 0, err
	}
	return num, nil

}

func TeamIsFull(num int, teamid string) (isfull bool, err error) {
	isfull, err = dao.TeamIsFull(num, teamid)
	if err != nil {
		return false, err
	}
	return isfull, nil
}

func NewTeam(newTeam *model.Team) error {
	err := dao.NewTeam(newTeam)
	if err != nil {
		return err
	}
	return nil
}

func CheckTeamIdExist(s string) (bool, error) {
	isExist, err := dao.CheckTeamIdExist(s)
	if err != nil {
		return isExist, err
	}
	return isExist, nil

}

func FetchTeamList() (TeamList []model.Team, err error) {
	TeamList, err = dao.FetchTeamList()
	if err != nil {
		return nil, err
	}
	return TeamList, nil

}

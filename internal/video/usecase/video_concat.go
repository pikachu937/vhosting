package usecase

import "github.com/dmitrij/vhosting/internal/video"

func (u *VideoUseCase) GetNonconcatedPaths(id int) (*[]video.NonCatVideo, error) {
	return u.videoRepo.GetNonconcatedPaths(id)
}

func (u *VideoUseCase) GetVideoPaths(pathStream, startDatetime string, durationRecord int) (*[]string, error) {
	return u.videoRepo.GetVideoPaths(pathStream, startDatetime, durationRecord)
}

func (u *VideoUseCase) UpdateReqVidArcFields(id, recordStatus, recordSize int) error {
	return u.videoRepo.UpdateReqVidArcFields(id, recordStatus, recordSize)
}

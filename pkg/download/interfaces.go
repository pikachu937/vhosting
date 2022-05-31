package download

type DownloadUseCase interface {
	IsValidExtension(file_name string) bool
	CreateDownloadLink(local_file_path string) *Download
}

package cmd

func HandleFileFormat(convertType string) string {
	switch convertType {
	case "pdf":
		return "application/pdf"
	case "odt":
		return "application/vnd.oasis.opendocument.text"
	case "docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case "txt":
		return "text/plain"
	case "doc":
		return "application/msword"
	default:
		return "application/octet-stream"
	}
}

func HandleImageFormat(convertType string) string {
	const jpeg = "image/jpeg"

	switch convertType {
	case "png":
		return "image/png"
	case "jpg":
		return jpeg
	case "jpeg":
		return jpeg
	case "gif":
		return "image/gif"
	case "bmp":
		return "image/bmp"
	case "webp":
		return "image/webp"
	default:
		return "application/octet-stream"
	}
}

make:
	@go run pixelate.go

dimensions:
	@ffprobe -v error -show_entries stream=width,height -of csv=p=0:s=x bliss-pixelated.jpg

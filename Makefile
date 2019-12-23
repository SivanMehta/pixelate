dimensions:
	@ffprobe -v error -show_entries stream=width,height -of csv=p=0:s=x frames/bliss-4k.jpg

frames/*.jpg:
	@go run main.go

output.gif: frames/*.jpg
	@echo Concatenating frames into a .gif ...
	@convert -delay 100 -loop 0 bliss-4k.jpg frames/bliss-{20,40,48,60,80,120,240}.jpg output.gif
	@rm frames/*.jpg

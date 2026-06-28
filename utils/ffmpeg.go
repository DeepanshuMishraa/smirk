package utils

import (
	"fmt"
	"os/exec"
)

type VideoOutput struct {
	Video360  string
	Video480  string
	Video720  string
	Video1080 string
}

func TranscodeVideo(videoPath string) (*VideoOutput, error) {

	resolutions := []int{360, 480, 720, 1080}

	outputs := make(map[int]string)

	for _, resolution := range resolutions {

		outputPath := fmt.Sprintf("/tmp/vid-pro.%dp.mp4", resolution)

		cmd := exec.Command(
			"ffmpeg",
			"-i", videoPath,
			"-vf", fmt.Sprintf("scale=-2:%d", resolution),
			"-c:v", "libx264",
			"-crf", "18",
			"-preset", "veryfast",
			"-c:a", "aac",
			outputPath,
		)

		output, err := cmd.CombinedOutput()
		if err != nil {
			return nil, fmt.Errorf(
				"failed to transcode %dp video: %v\n%s",
				resolution,
				err,
				string(output),
			)
		}

		outputs[resolution] = outputPath
	}

	return &VideoOutput{
		Video360:  outputs[360],
		Video480:  outputs[480],
		Video720:  outputs[720],
		Video1080: outputs[1080],
	}, nil
}


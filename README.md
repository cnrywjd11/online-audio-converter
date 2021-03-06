# online_audio_converter
간단한 온라인 오디오 변환기 (서버)

## build
install dependency library, `ffmpeg` for github.com/FFmpeg/FFmpeg:

    $ sudo apt install ffmpeg # Linux
    $ brew install ffmpeg # Mac
    
server build:

    $ go build -o online_audio_converter_server cmd/main.go

## usage
run:

    $ ./online_audio_converter_server
## required header
- Accept-Audio-Format: base64_encoded_json
  - base64_encode({"codec":"mp3","samplingrate":44100,"channel":2,"bitrate":"96k","speed":"130%","volume":"3dB"})

## sample request
wav input -> 44k, stereo, 96k, 130% speed, +3dB, mp3 output:
```curl
curl -v -X POST 'localhost/convert' \
-H 'Content-Type: multipart/form-data' \
-H 'Accept-Audio-Format: eyJjb2RlYyI6Im1wMyIsInNhbXBsaW5ncmF0ZSI6NDQxMDAsImNoYW5uZWwiOjIsImJpdHJhdGUiOiI5NmsiLCJzcGVlZCI6IjEzMCUiLCJ2b2x1bWUiOiIzZEIifQo=' \
-F 'audio=@sample.wav'
```
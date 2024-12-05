resp=$(curl -X POST -d "{\"voiceId\":\"zh-CN-XiaoyiNeural\",\"speed\":\"+40%\",\"text\":\"$1\"}" http://127.0.0.1:19020/generalMp3)
echo $resp | jq '.output.mp3_hex' | xxd -r -p > $2.mp3
echo $resp | jq '.output.subtitles_hex' | xxd -r -p > $2.vvt
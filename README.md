# mp3-analyzer

Right now, this program takes an MP3 file piped via STDIN, analyzes it for ID3 information as well as length and size, and dumps them to STDOUT. To finish, this thing needs to do more, like analyze the bitrate of an MP3 file dynamically so the size calculation is valid for all bitrates. Also, it should be more resilient to errors.
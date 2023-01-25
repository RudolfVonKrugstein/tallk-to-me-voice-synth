#!/bin/bash

# list all voices
ALL_VOICES=

for file in $(find text -name lines.txt); do
  VOICE=$(echo $file | sed 's|text/||g' | sed 's|/lines.txt||g')
  NEURAL_PARAMETER=""
  if [[ $VOICE == *_neural ]]
  then
    NEURAL_PARAMETER="--engine neural"
    VOICE=$(echo $VOICE | sed 's|_neural||g')
  fi
  echo $VOICE

  PREFIX_FILE=$(echo $file | sed 's|lines.txt|prefix.txt|g')
  PREFIX=$(cat $PREFIX_FILE)
  PREFIX=$(echo $PREFIX)

  echo "Handling lines for $VOICE"

  cat $file | while read line
  do
    echo "LINE: $line"
    echo "Voice Parameters: $VOICE $NEURAL_PARAMETER"
    aws polly synthesize-speech \
    --output-format mp3 \
    --voice-id $VOICE $NEURAL_PARAMETER \
    --text "$line" \
    "${PREFIX}_${VOICE}_$(echo $line | sed 's| |_|g' ).mp3"
  done
done



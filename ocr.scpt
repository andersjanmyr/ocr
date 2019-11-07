-- ~/Library/Scripts/Folder Action Scripts/ocr.scpt
on adding folder items to theAttachedFolder after receiving theNewItems
  tell application "Finder"
    repeat with anItem in theNewItems
      set AppleScript's text item delimiters to {return}
      set p to POSIX path of anItem
      set command to "GOOGLE_APPLICATION_CREDENTIALS=/path/to/auth.json /usr/local/bin/ocr " & (quoted form of p)
      try
	set output to do shell script command
	set the clipboard to output
      on error theError
	activate
	display dialog theError
      end try
    end repeat
  end tell
end adding folder items to


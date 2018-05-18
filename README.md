# OCR

`ocr` uses GCP or AWS to read text from images.


## Setup

* For Google, enable the vision API and billing and create a service account, https://cloud.google.com/vision/docs/libraries#client-libraries-install-go, then make sure that you set `GOOGLE_APPLICATION_CREDENTIALS` to the path to your new service account authentication file.
* For AWS, configure your credentials according to https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html

## Usage

```
ocr <image file>
```

## Folder Watching

The following AppleScript can be used to listen for new files and copy the text
of them into the clipboard

```
on adding folder items to theAttachedFolder after receiving theNewItems
	tell application "Finder"
		
		repeat with anItem in theNewItems
			set p to POSIX path of anItem
			
			set command to "GOOGLE_APPLICATION_CREDENTIALS=/Users/vn0wxf9/gocode/src/github.com/andersjanmyr/ocr/ocr-sa.json /Users/vn0wxf9/gocode/src/github.com/andersjanmyr/ocr/ocr " & (quoted form of p)
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
```

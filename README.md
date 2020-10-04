# Image-tag-tracker
Adds image tags to txt file and get the latest tag.

# Compile
```
go install
go build -o main.out
```

# Usage
Get the latest tag
```
./main.out -fpath=tags.txt tags latest
```

Add a new tag
```
./main.out -fpath=tags.txt tags add your-tag-to-add-here
```
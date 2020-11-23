# IMAGE PROCESSING (PROJECT)

- Work with images in go
  https://github.com/disintegration/imaging

`` go get -u github.com/disintegration/imaging ``


# Filter package 
`` filter/filter.go ``

``package filter``

``type Filter interface {``
    ``Process(srcPath, dstPath string) error``
``}``
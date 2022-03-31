# golang-image-ico

golang support for windows .ico file format

> http://en.wikipedia.org/wiki/ICO_(file_format)

## Icon resource structure

All values in ICO/CUR files are represented in **[little-endian]** byte order.

### ICONDIR structure

| Offset# | Size | Purpose                                                                                             |
|---------|------|-----------------------------------------------------------------------------------------------------|
| 0       | 2    | Reserved. Must always be 0.                                                                         |
| 2       | 2    | Specifies image type: 1 for icon (.ICO) image, 2 for cursor (.CUR) image. Other values are invalid. |
| 4       | 2    | Specifies number of images in the file.                                                             |

### Structure of image directory

|            |                            |
|------------|----------------------------|
| Image #1   | Entry for the first image  |
| Image #2   | Entry for the second image |
| Image #n   | Entry for the last image   |

### ICONDIRENTRY structure

| Offset# | Size | Purpose                                                                                                                                                            |
|---------|------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 0       | 1    | Specifies image width in pixels. Can be any number between 0 and 255. Value 0 means image width is 256 pixels.                                                     |
| 1       | 1    | Specifies image height in pixels. Can be any number between 0 and 255. Value 0 means image height is 256 pixels.                                                   |
| 2       | 1    | Specifies number of colors in the color palette. Should be 0 if the image does not use a color palette.                                                            |
| 3       | 1    | Reserved. Should be 0.                                                                                                                                             |
| 4       | 2    | ● In ICO format: Specifies color planes. Should be 0 or 1. ● In CUR format: Specifies the horizontal coordinates of the hotspot in number of pixels from the left. |
| 6       | 2    | ● In ICO format: Specifies bits per pixel. ●In CUR format: Specifies the vertical coordinates of the hotspot in number of pixels from the top.                     |
| 8       | 4    | Specifies the size of the image's data in bytes                                                                                                                    |
| 12      | 4    | Specifies the offset of BMP or PNG data from the beginning of the ICO/CUR file                                                                                     |

[little-endian]: https://en.wikipedia.org/wiki/Endianness

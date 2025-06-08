#This program should extract all

#Options to handle:
# ./spider [-rlp] URL
# -r recursively downloads the images in a URL received as paramater
# -r -l [N] indicates the maximum  depth level of the recursion, 5 if not indicated
# -p [PATH] indicates the path where the downloaded files will be saved, if not specified ./data/ will be used

#The program will download the following extensions by default:
# .jpg/jpeg
# .png 
# .gif
# .bmp

#Modules to use: argparse, requests, BeautifulSoup, os, pathlib, shutil

import sys

r_mask = 0xb1000
l_mask = 0xb0100
p_mask = 0xb0010

def parse_args():
    flags = 0b0000
    if sys.argv[1] == '-r':
        return (flag & r_mask)

def main():
    if len(sys.argv) > 1:
        return ;
    flags = parse_args()
    print(f"{flag:08b}")
    if __name__ == "__main__":
        main()


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

r_mask = 0b1000
l_mask = 0b0100
p_mask = 0b0010

def parse_args():
    flags = 0b0000
    for arg in sys.argv[1:]:
        if arg.startswith('-'):
            for c in arg[1:]:
                if c == 'r':
                    flags |= r_mask
                elif c == 'p':
                    flags |= p_mask
                elif c == 'l':
                    flags |= l_mask
                else:
                    print(f"Unknown Option: {c}")
    return flags

# def spider_exec():


def main():
    flags = parse_args()
    if flags == 0:
        return
    print(f"{flags:03b}")


if __name__ == "__main__":
    main()

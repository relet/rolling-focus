enfuse -o $1.jpg \
 --exposure-weight=0 \
 --saturation-weight=0 \
 --contrast-weight=1 \
 --hard-mask \
 --gray-projector=l-star \
 --contrast-edge-scale=0.3 \
 $1/*.jpg


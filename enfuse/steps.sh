DEPTH=20
NUMFILES=$(ls $1/*.jpg | wc -l)
FOLDERS=$(( $NUMFILES / $DEPTH + 1 ))

for i in $(seq 1 $FOLDERS); do
  mkdir $1/$i
  FILES=$(ls $1/*.jpg | head -n $DEPTH)
  mv $FILES $1/$i
done

for i in $(seq 1 $FOLDERS); do
  ./auto.sh $1/$i
done

./auto.sh $1

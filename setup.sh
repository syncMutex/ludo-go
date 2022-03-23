if [[ "$1" != "" && "$1" != "-noserver" ]]; then
	echo error: invalid args
	exit 1
fi

if [ "$1" == "" ]; then
	cd server
	pip install pyngrok
	cd ../
fi

cd ./game
echo building game binary...
go build

mv ludo* ../

echo Setup successful!

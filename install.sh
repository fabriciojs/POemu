echo "Installing...";

if [ ! -d "bin" ]; then
	mkdir bin;
	chmod 777 bin;
	echo "Directory bin created.";
else
	echo "Directoy bin already exists.";
fi


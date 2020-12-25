run: jar

jar: classes
	jar cfm Juego.jar manifest.mf -C build/classes .


classes: dir
	javac -sourcepath src/ -d build/classes src/Tarea3/*.java

dir:
	if [ ! -d build/classes ]; then mkdir -p build/classes; fi
	if [ ! -d build/jar ]; then mkdir -p build/jar; fi

clean:
	if [ -d build ]; then rm build -R; fi

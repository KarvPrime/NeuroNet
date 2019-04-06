

SET "NoNeuroData=true"



IF NOT EXIST "%~dp0NeuroNet\Data\Train\mnist_train.csv" (

	SET "NoNeuroData="

	ECHO No training data for NeuroNet.

)


IF NOT EXIST "%~dp0NeuroNet\Data\Test\mnist_test.csv" (

	SET "NoNeuroData="

	ECHO No testing data for NeuroNet.

)



IF NOT DEFINED NoNeuroData (

	SET "NoTranslateData=true"
	
IF NOT EXIST "%~dp0MNISTTranslate\mnist_train.csv" (

		SET "NoTranslateData="
		ECHO No translated training data.

	)

	
IF NOT EXIST "%~dp0MNISTTranslate\mnist_test.csv" (

		SET "NoTranslateData="

		ECHO No translated testing data.

	)

	

	IF NOT DEFINED NoTranslateData (

		cd %~dp0MNISTTranslate\

		IF NOT EXIST "MNISTTranslate.exe" (

			ECHO No translator build. Building...
			go build MNISTTranslate.go

			ECHO Building complete.

		)

		ECHO Starting translation.

		start MNISTTranslate.exe

		cd ..

		ECHO Translation complete.
	)

	

	ECHO Moving data.
	xcopy %~dp0MNISTTranslate\mnist_train.csv %~dp0NeuroNet\Data\Train\mnist_train.csv

	xcopy %~dp0MNISTTranslate\mnist_test.csv %~dp0NeuroNet\Data\Test\mnist_test.csv

)



cd ./NeuroNet

IF NOT EXIST "NeuroNet.exe" (

	ECHO No NeuroNet build. Building...
	go build NeuroNet.go

	ECHO Building complete.
)

ECHO Starting NeuroNet.
ECHO __________________
start NeuroNet.exe

cd ..

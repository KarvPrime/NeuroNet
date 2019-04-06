# NeuroNet
Go Concurrent Neural Network

## How to run

### Preparation
* Put MNISTTranslate, NeuroNet, and the Bash- and/or Batchfile into your GOROOT directory
* Download the [MNIST Dataset](http://yann.lecun.com/exdb/mnist/)
* Put the files into the ./MNISTTranslate directory

### Quick run
* Run the program
	* Linux: ./BakkBash
	* Windows: ./BakkBatch.cmd

### Custom
* Change ./NeuroNet/Batchfile/Batch to your specifications
* Modify how your network should look like (example in ./NeuroNet/Data/Networks/Numbers)
* First time (or after changes) you need to build it
	* MNISTTranslate: go build ./MNISTTranslate/MNISTTranslate.go
	* NeuroNet: go build ./NeuroNet/NeuroNet.go
* Run MNISTTranslate
* Put the translated files into your specified folders
	* Standard Test folder: ./NeuroNet/Data/Test
	* Standard Train folder: ./NeuroNet/Data/Train
* Run NeuroNet

## Batchfile
Each line in the batch file represents one action which the software will
execute.
* ResultFile:<Filepath> → Path to the result output
* NetworkFile:<Filepath> → Path to the network creation file
* PersistenceFile:<Filepath> → Path to the persistence
* TrainFile:<Filepath> → Path to the training data
* TestFile:<Filepath> → Path to the test data
* PreProcessing:<PreProcessing> → None / MeanSubstraction / Proportional
* Parallel:<Integer> → Number of goroutines
* WorkerBatch:<Integer> → Batch size for each worker before result merge
* LearningRate:<Float> → Rate of learning
* Lambda:<Float> → Lambda value for elastic net regularization
* MinWeight:<Float> → Minimum weight for connectome initialization
* MaxWeight:<Float> → Maximum weight for connectome initialization
* TargetColumnsStart:<Integer> → First field that is a target value
* TargetColumnsEnd:<Integer> → First field that isn’t a target value
* Train:<Integer> → Train the network x times
* Test:<Integer> → Run a test x times


## Network
Each line represents one layer in the network.

Schema: Activation,Neurons

Current activation functions:
* Identity
* Logistic
* TanH
* ReLU
* LeakyReLU
* ELU
* SoftMax

In example "SoftMax, 10" (without quotation marks) creates one SoftMax
layer with 10 neurons. Currently a bias neuron will be added to every but
the last layer.
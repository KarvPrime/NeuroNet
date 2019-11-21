options(scipen=9001)

#Lenovo
lenovoCore1 = read.csv(file="./Data/Final/Result/Lenovo/Results-100-1", sep=";", stringsAsFactors=F, na.strings="#NV")
lenovoCore2 = read.csv(file="./Data/Final/Result/Lenovo/Results-100-2", sep=";", stringsAsFactors=F, na.strings="#NV")
lenovoCore3 = read.csv(file="./Data/Final/Result/Lenovo/Results-100-3", sep=";", stringsAsFactors=F, na.strings="#NV")
lenovoCore4 = read.csv(file="./Data/Final/Result/Lenovo/Results-100-4", sep=";", stringsAsFactors=F, na.strings="#NV")
lenovoCore5 = read.csv(file="./Data/Final/Result/Lenovo/Results-100-5", sep=";", stringsAsFactors=F, na.strings="#NV")
lenovoCore6 = read.csv(file="./Data/Final/Result/Lenovo/Results-100-6", sep=";", stringsAsFactors=F, na.strings="#NV")
lenovoCore7 = read.csv(file="./Data/Final/Result/Lenovo/Results-100-7", sep=";", stringsAsFactors=F, na.strings="#NV")
lenovoCore8 = read.csv(file="./Data/Final/Result/Lenovo/Results-100-8", sep=";", stringsAsFactors=F, na.strings="#NV")

lenovoCores = rbind(lenovoCore1, lenovoCore2, lenovoCore3, lenovoCore4, lenovoCore5, lenovoCore6, lenovoCore7, lenovoCore8)
lenovoTrain = subset(lenovoCores, lenovoCores$Mode %in% c("train"))
lenovoTest = subset(lenovoCores, lenovoCores$Mode %in% c("test"))

lenovoTest <- transform(lenovoTest, Percent = Correct / EpochLines * 100)

lenovoTrainSubset <- subset(lenovoTest, lenovoTest$EpochLines %in% 60000)
lenovoTestSubset <- subset(lenovoTest, lenovoTest$EpochLines %in% 10000)
lenovoTrainMax = subset(lenovoTrainSubset, lenovoTrainSubset$TotalLines %in% 6000000)
lenovoTestMax = subset(lenovoTestSubset, lenovoTestSubset$TotalLines %in% 6000000)
lenovoMaxTestTrain = merge(lenovoTrainMax, lenovoTestMax[, c("Parallel", "Percent")], by="Parallel")

jpeg("./Data/Final/Images/lenovoMaxTestTrain.jpg", width=800, height=1000)
matplot(cbind(lenovoMaxTestTrain[13], lenovoMaxTestTrain[14]), type=c("b"), pch=1, col=1:10, ylim=c(96, 100), main="Training and test set accuracy", sub="Tested with Lenovo Yoga 2", xlab="Goroutines", ylab="Accuracy")
legend("bottomleft", legend=c("Train", "Test"), col=1:2, pch=1)
dev.off()

lenovoTrain1 <- subset(lenovoTrainSubset, lenovoTrainSubset$Parallel %in% 1)
lenovoTest1 <- subset(lenovoTestSubset, lenovoTestSubset$Parallel %in% 1)

jpeg("./Data/Final/Images/lenovoAccuracy1.jpg", width=800, height=1000)
matplot(cbind(lenovoTrain1$Percent, lenovoTest1$Percent), type=c("b"), pch=1, col=1:2, ylim=c(96, 100), main="Training and test set accuracy (1 goroutine)", sub="Tested with Lenovo Yoga 2",  xlab="Epoch (*10)", ylab="Accuracy")
legend("bottomright", legend=c("Train", "Test"), col=1:2, pch=1)
dev.off()

lenovoTrain2 <- subset(lenovoTrainSubset, lenovoTrainSubset$Parallel %in% 2)
lenovoTest2 <- subset(lenovoTestSubset, lenovoTestSubset$Parallel %in% 2)

jpeg("./Data/Final/Images/lenovoAccuracy2.jpg", width=800, height=1000)
matplot(cbind(lenovoTrain2$Percent, lenovoTest2$Percent), type=c("b"), pch=1, col=1:2, ylim=c(96, 100), main="Training and test set accuracy (2 goroutines)", sub="Tested with Lenovo Yoga 2",  xlab="Epoch (*10)", ylab="Accuracy")
legend("bottomright", legend=c("Train", "Test"), col=1:2, pch=1)
dev.off()

lenovoTrain4 <- subset(lenovoTrainSubset, lenovoTrainSubset$Parallel %in% 4)
lenovoTest4 <- subset(lenovoTestSubset, lenovoTestSubset$Parallel %in% 4)

jpeg("./Data/Final/Images/lenovoAccuracy4.jpg", width=800, height=1000)
matplot(cbind(lenovoTrain4$Percent, lenovoTest4$Percent), type=c("b"), pch=1, col=1:2, ylim=c(96, 100), main="Training and test set accuracy (4 goroutines)", sub="Tested with Lenovo Yoga 2",  xlab="Epoch (*10)", ylab="Accuracy")
legend("bottomright", legend=c("Train", "Test"), col=1:2, pch=1)
dev.off()

lenovoTrain8 <- subset(lenovoTrainSubset, lenovoTrainSubset$Parallel %in% 8)
lenovoTest8 <- subset(lenovoTestSubset, lenovoTestSubset$Parallel %in% 8)

jpeg("./Data/Final/Images/lenovoAccuracy8.jpg", width=800, height=1000)
matplot(cbind(lenovoTrain8$Percent, lenovoTest8$Percent), type=c("b"), pch=1, col=1:2, ylim=c(96, 100), main="Training and test set accuracy (8 goroutines)", sub="Tested with Lenovo Yoga 2",  xlab="Epoch (*10)", ylab="Accuracy")
legend("bottomright", legend=c("Train", "Test"), col=1:2, pch=1)
dev.off()

lenovoTotalTimes = aggregate(TotalElapsedTime~Parallel, data=lenovoTrain, max)

jpeg("./Data/Final/Images/lenovoAllGoroutines.jpg", width=800, height=1000)
boxplot(split(lenovoTrain$EpochElapsedTime, lenovoTrain$Parallel), main="Computation speed (Worker batch size of 100)", sub="Tested with Lenovo Yoga 2", xlab="Worker Goroutines (100 epochs each)", ylab="Time in seconds (for each epoch)", notch=TRUE)
dev.off()

lenovoCore1train = subset(lenovoCore1, lenovoCore1$Mode %in% c("train"))
min(lenovoCore1train[,7])
max(lenovoCore1train[,7])
median(lenovoCore1train[,7])

lenovoCore2train = subset(lenovoCore2, lenovoCore2$Mode %in% c("train"))
min(lenovoCore2train[,7])
max(lenovoCore2train[,7])
median(lenovoCore2train[,7])

lenovoCore4train = subset(lenovoCore4, lenovoCore4$Mode %in% c("train"))
min(lenovoCore4train[,7])
max(lenovoCore4train[,7])
median(lenovoCore4train[,7])

lenovoCore8train = subset(lenovoCore8, lenovoCore8$Mode %in% c("train"))
min(lenovoCore8train[,7])
max(lenovoCore8train[,7])
median(lenovoCore8train[,7])


#BPI
bpiCore1 = read.csv(file="./Data/Final/Result/BPI/FullResults-1", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore2 = read.csv(file="./Data/Final/Result/BPI/FullResults-2", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore3 = read.csv(file="./Data/Final/Result/BPI/FullResults-3", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore4 = read.csv(file="./Data/Final/Result/BPI/FullResults-4", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore5 = read.csv(file="./Data/Final/Result/BPI/FullResults-5", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore6 = read.csv(file="./Data/Final/Result/BPI/FullResults-6", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore7 = read.csv(file="./Data/Final/Result/BPI/FullResults-7", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore8 = read.csv(file="./Data/Final/Result/BPI/FullResults-8", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore9 = read.csv(file="./Data/Final/Result/BPI/FullResults-9", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore10 = read.csv(file="./Data/Final/Result/BPI/FullResults-10", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore11 = read.csv(file="./Data/Final/Result/BPI/FullResults-11", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore12 = read.csv(file="./Data/Final/Result/BPI/FullResults-12", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore13 = read.csv(file="./Data/Final/Result/BPI/FullResults-13", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore14 = read.csv(file="./Data/Final/Result/BPI/FullResults-14", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore15 = read.csv(file="./Data/Final/Result/BPI/FullResults-15", sep=";", stringsAsFactors=F, na.strings="#NV")
bpiCore16 = read.csv(file="./Data/Final/Result/BPI/FullResults-16", sep=";", stringsAsFactors=F, na.strings="#NV")

bpiCores = rbind(bpiCore1, bpiCore2, bpiCore3, bpiCore4, bpiCore5, bpiCore6, bpiCore7, bpiCore8, bpiCore9, bpiCore10, bpiCore11, bpiCore12, bpiCore13, bpiCore14, bpiCore15, bpiCore16)
bpiTrain = subset(bpiCores, bpiCores$Mode %in% c("train"))
bpiTest = subset(bpiCores, bpiCores$Mode %in% c("test"))

bpiTest <- transform(bpiTest, Percent = Correct / EpochLines * 100)

bpiTrainSubset <- subset(bpiTest, bpiTest$EpochLines %in% 60000)
bpiTestSubset <- subset(bpiTest, bpiTest$EpochLines %in% 10000)
bpiTrainMax = subset(bpiTrainSubset, bpiTrainSubset$TotalLines %in% 6000000)
bpiTestMax = subset(bpiTestSubset, bpiTestSubset$TotalLines %in% 6000000)
bpiMaxTestTrain = merge(bpiTrainMax, bpiTestMax[, c("Parallel", "Percent")], by="Parallel")

jpeg("./Data/Final/Images/bpiMaxTestTrain.jpg", width=800, height=1000)
matplot(cbind(bpiMaxTestTrain[13], bpiMaxTestTrain[14]), type=c("b"), pch=1, col=1:10, ylim=c(96, 100), main="Training and test set accuracy", sub="Tested with Banana Pi M3", xlab="Goroutines", ylab="Accuracy")
legend("bottomright", legend=c("Train", "Test"), col=1:2, pch=1)
dev.off()

bpiTrain1 <- subset(bpiTrainSubset, bpiTrainSubset$Parallel %in% 1)
bpiTest1 <- subset(bpiTestSubset, bpiTestSubset$Parallel %in% 1)

jpeg("./Data/Final/Images/bpiAccuracy1.jpg", width=800, height=1000)
matplot(cbind(bpiTrain1$Percent, bpiTest1$Percent), type=c("b"), pch=1, col=1:2, ylim=c(80, 100), main="Training and test set accuracy (1 goroutine)", sub="Tested with Banana Pi M3",  xlab="Epoch", ylab="Accuracy")
legend("bottomright", legend=c("Train 1 Goroutine", "Test 1 Goroutine"), col=1:2, pch=1)
dev.off()

bpiTrain2 <- subset(bpiTrainSubset, bpiTrainSubset$Parallel %in% 2)
bpiTest2 <- subset(bpiTestSubset, bpiTestSubset$Parallel %in% 2)

jpeg("./Data/Final/Images/bpiAccuracy2.jpg", width=800, height=1000)
matplot(cbind(bpiTrain2$Percent, bpiTest2$Percent), type=c("b"), pch=1, col=1:2, ylim=c(80, 100), main="Training and test set accuracy (2 goroutines)", sub="Tested with Banana Pi M3",  xlab="Epoch", ylab="Accuracy")
legend("bottomright", legend=c("Train 2 Goroutines", "Test 2 Goroutines"), col=1:2, pch=1)
dev.off()

bpiTrain4 <- subset(bpiTrainSubset, bpiTrainSubset$Parallel %in% 4)
bpiTest4 <- subset(bpiTestSubset, bpiTestSubset$Parallel %in% 4)

jpeg("./Data/Final/Images/bpiAccuracy4.jpg", width=800, height=1000)
matplot(cbind(bpiTrain4$Percent, bpiTest4$Percent), type=c("b"), pch=1, col=1:2, ylim=c(80, 100), main="Training and test set accuracy (4 goroutines)", sub="Tested with Banana Pi M3",  xlab="Epoch", ylab="Accuracy")
legend("bottomright", legend=c("Train 4 Goroutines", "Test 4 Goroutines"), col=1:2, pch=1)
dev.off()

bpiTrain8 <- subset(bpiTrainSubset, bpiTrainSubset$Parallel %in% 8)
bpiTest8 <- subset(bpiTestSubset, bpiTestSubset$Parallel %in% 8)

jpeg("./Data/Final/Images/bpiAccuracy8.jpg", width=800, height=1000)
matplot(cbind(bpiTrain8$Percent, bpiTest8$Percent), type=c("b"), pch=1, col=1:2, ylim=c(80, 100), main="Training and test set accuracy (8 goroutines)", sub="Tested with Banana Pi M3",  xlab="Epoch", ylab="Accuracy")
legend("bottomright", legend=c("Train 8 Goroutines", "Test 8 Goroutines"), col=1:2, pch=1)
dev.off()

bpiTrain12 <- subset(bpiTrainSubset, bpiTrainSubset$Parallel %in% 12)
bpiTest12 <- subset(bpiTestSubset, bpiTestSubset$Parallel %in% 12)

jpeg("./Data/Final/Images/bpiAccuracy12.jpg", width=800, height=1000)
matplot(cbind(bpiTrain12$Percent, bpiTest12$Percent), type=c("b"), pch=1, col=1:2, ylim=c(80, 100), main="Training and test set accuracy (12 goroutines)", sub="Tested with Banana Pi M3",  xlab="Epoch", ylab="Accuracy")
legend("bottomright", legend=c("Train 12 Goroutines", "Test 12 Goroutines"), col=1:2, pch=1)
dev.off()

bpiTrain16 <- subset(bpiTrainSubset, bpiTrainSubset$Parallel %in% 16)
bpiTest16 <- subset(bpiTestSubset, bpiTestSubset$Parallel %in% 16)

jpeg("./Data/Final/Images/bpiAccuracy16.jpg", width=800, height=1000)
matplot(cbind(bpiTrain16$Percent, bpiTest16$Percent), type=c("b"), pch=1, col=1:2, ylim=c(80, 100), main="Training and test set accuracy (16 goroutines)", sub="Tested with Banana Pi M3",  xlab="Epoch", ylab="Accuracy")
legend("bottomright", legend=c("Train 16 Goroutines", "Test 16 Goroutines"), col=1:2, pch=1)
dev.off()

bpiTotalTimes=aggregate(TotalElapsedTime~Parallel, data=bpiTrain, max)

jpeg("./Data/Final/Images/bpiAllGoroutines.jpg", width=800, height=1000)
boxplot(split(bpiTrain$EpochElapsedTime, bpiTrain$Parallel), main="Computation speed (Worker batch size of 100)", sub="Tested with Banana Pi M3", xlab="Worker Goroutines (100 epochs each)", ylab="Time in seconds (for each epoch)", notch=TRUE)
dev.off()

bpiCore1train = subset(bpiCore1, bpiCore1$Mode %in% c("train"))
min(bpiCore1train[,7])
max(bpiCore1train[,7])
median(bpiCore1train[,7])

bpiCore2train = subset(bpiCore2, bpiCore2$Mode %in% c("train"))
min(bpiCore2train[,7])
max(bpiCore2train[,7])
median(bpiCore2train[,7])

bpiCore4train = subset(bpiCore4, bpiCore4$Mode %in% c("train"))
min(bpiCore4train[,7])
max(bpiCore4train[,7])
median(bpiCore4train[,7])

bpiCore8train = subset(bpiCore8, bpiCore8$Mode %in% c("train"))
min(bpiCore8train[,7])
max(bpiCore8train[,7])
median(bpiCore8train[,7])

bpiCore12train = subset(bpiCore12, bpiCore12$Mode %in% c("train"))
min(bpiCore12train[,7])
max(bpiCore12train[,7])
median(bpiCore12train[,7])

bpiCore16train = subset(bpiCore16, bpiCore16$Mode %in% c("train"))
min(bpiCore16train[,7])
max(bpiCore16train[,7])
median(bpiCore16train[,7])


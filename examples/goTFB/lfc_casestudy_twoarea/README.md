# lfc_casestudy_twoarea

This example is based on the mathematics presented in Section 2.2 (_Automatic Load-Frequency Control in Two-Area Systems_)
of *Application of Neural Networks to Load-Frequency Control in Power Systems* in _Neural Networks_

Link: https://www.sciencedirect.com/science/article/pii/0893608094900671

In the original article they used an integral controller as the comparison against their feedforward neural networks.
We will compare against the integral controller, as well as their neural network.
There are other papers which use this example as the basis for their neural networks.

## USAGE NOTES

* Be careful how you tick the simulation and where delays are introduced.
* The "load" block must be ticked separately to the rest of the math logic due to the synchronous PRE

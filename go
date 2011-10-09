#!/usr/bin/env bash 
# Step 1: Build the Independent code
cd independent 
mvn  -s ../settings.xml clean deploy  
cd - 

# Step 2: Remove all the com.builddoctor code from the local repo (force it to go to S3
rm -rf ${HOME}/.m2/repository/com/builddoctor

# Step 3: Build the Dependent code
cd dependent 
mvn  -s ../settings.xml clean install
cd - 

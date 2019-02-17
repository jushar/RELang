@echo off

SET CLASSPATH=.;.\tools\antlr-4.7.2-complete.jar;%CLASSPATH%

java org.antlr.v4.Tool %*

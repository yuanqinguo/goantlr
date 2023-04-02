antlr4=java -Xmx500M -cp "/mnt/d/applications/antlr4/antlr-4.10.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool


all: generate
generate:
	$(antlr4) -Dlanguage=Go -no-listener -visitor -o ../parser express.g4
SRCS = srcs/lists.cpp srcs/main.cpp tasks/rotate.cpp

HEADERS = includes/listf.h

CFLAGS = -Wall -Werror -Wextra 

CC = g++

all : output

output :
		${CC}  ${HEADERS} ${SRCS} -o output

clean :
		rm -rf output *.gch
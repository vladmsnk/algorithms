SRCS = lists.cpp main.cpp reverse_list.cpp

HEADERS = listf.h

CFLAGS = -Wall -Werror -Wextra 

CC = g++

all : output

output :
		${CC} ${CFLAGS} ${HEADERS} ${SRCS} -o output

clean :
		rm -rf output *.gch
FROM busybox:glibc
ADD shopping /bin/shopping
EXPOSE 8080
CMD ["shopping"]
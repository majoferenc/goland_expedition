FROM golang:1.13.0-alpine3.10 AS build-env

# Allow Go to retrive the dependencies for the build step
RUN apk add --no-cache git

WORKDIR /awesomeProject/
ADD . /awesomeProject/

# Compile the binary, we don't want to run the cgo resolver
RUN CGO_ENABLED=0 go build -o /awesomeProject/app .

# final stage, extra small image 11 MB with our app!
FROM scratch

COPY --from=build-env /awesomeProject/app /

EXPOSE 8091
EXPOSE 4040

CMD ["/app"]
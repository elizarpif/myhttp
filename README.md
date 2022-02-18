# myhttp

A tool which makes http requests and prints the address of the request along with the
MD5 hash of the response.

• The tool able to perform the requests in parallel

• The tool able to limit the number of parallel requests, to prevent exhausting local resources.
The tool can accept a flag to indicate this limit, and it defaults to 10 if the flag is not provided.

## As info

1) Flags. The tool has flag `parallel`, it accepts only unsigned int value
2) Errors. If tool gets an error, only address will be printed
3) Timeout. Http-requests have a timeout of 5 seconds

### Example of work

```shell
./myhttp -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com baoooooooooooo.com
```
```
baoooooooooooo.com
adjust.com fe3dd0d3ace543820dcc378bc85efe42
facebook.com 3684c2e12b08e342c36c2eb61adf3410
yandex.com 8be5dc4e9f35e32ddbfc1bb59cfd973a
yahoo.com f39faafa626f5f052b0119fdebdbe927
twitter.com bf70ff5c890f44fda01f4873d68ec8e5
baroquemusiclibrary.com dadd85ea2820b3fea639660ddc055c6e
reddit.com/r/funny ec9b8ebbf4258662204f581648fb68c7
google.com 47727d047af21c0a2ea96c58e117a553
reddit.com/r/notfunny 94a81d4cb09b92b3a74c4b8c87e6f5d1
```
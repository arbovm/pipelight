## pipelight

#### 'cause it shows the difference

### Install

    go get github.com/arbovm/pipelight
    
### Example

    » date | pipelight

> Fri Mar  1 18:53:49 CET 2013

It shows the date, awesome. Let's run it again:

    » date | pipelight
    
> Fri Mar  1 18:53:<b>54</b> CET 2013

It highlights the seconds of the date output because they changed since the last time the command was executed.
OK, that doesn't make any sense. But what about this:
  
Clear ``pipelight`` buffer

    echo -n | pipelight

And then take a look at the following:

                                            
    curl -Is https://github.com -c cookies.txt -b cookies.txt| pipelight
    
> HTTP/1.1 200 OK<br>
> Server: GitHub.com<br>
> Date: Fri, 01 Mar 2013 18:26:17 GMT<br>
> Content-Type: text/html; charset=utf-8<br>
> Connection: keep-alive<br>
> Status: 200 OK<br>
> Cache-Control: private, max-age=0, must-revalidate<br>
> Strict-Transport-Security: max-age=2592000<br>
> X-Frame-Options: deny<br>
> Set-Cookie: logged_in=no; domain=.github.com; path=/; expires=Tue, 01-Mar-2033 18:26:17 GMT<br>
> Set-Cookie: _gh_sess=BAh7Bz...d25af3; path=/; expires=Sun, 01-Jan-2023 00:00:00 GMT; secure; HttpOnly<br>
> X-Runtime: 6<br>
> ETag: "61b566ab19f54e54bf0dcee24a3dac19"<br>
> Content-Length: 10925<br>


And execute it again:


    curl -Is https://github.com -c cookies.txt -b cookies.txt| pipelight
    
> HTTP/1.1 200 OK<br>
> Server: GitHub.com<br>
> Date: Fri, 01 Mar 2013 18:26:<b>21</b> GMT<br>
> Content-Type: text/html; charset=utf-8<br>
> Connection: keep-alive<br>
> Status: 200 OK<br>
> Cache-Control: private, max-age=0, must-revalidate<br>
> Strict-Transport-Security: max-age=2592000<br>
> X-Frame-Options: deny<br>
> Set-Cookie: logged_in=no; domain=.github.com; path=/; expires=Tue, 01-Mar-2033 18:26:21 GMT<br>
> Set-Cookie: _gh_sess=BAh7Bz...d25af3; path=/; expires=Sun, 01-Jan-2023 00:00:00 GMT; secure; HttpOnly<br>
> X-Runtime: <b>7</b><br>
> ETag: "<b>2076</b>6<b>f7927ccf56</b>4<b>3ba845ac381917ab</b>"<br>
> Content-Length: 10925<br>


Now it shows, that only ``X-Runtime``, ``Date`` and parts of the ``ETag`` changed.

Let'a try this again using ``Accept-Encoding: gzip`` to force compresssion:

    curl -Is https://github.com -c cookies.txt -b cookies.txt -H'Accept-Encoding: gzip' | pipelight
    

> HTTP/1.1 200 OK<br>
> Server: GitHub.com<br>
> Date: Fri, 01 Mar 2013 18:26:<b>4</b>1 GMT<br>
> Content-Type: text/html; charset=utf-8<br>
> Connection: keep-alive<br>
> Status: 200 OK<br>
> Cache-Control: private, max-age=0, must-revalidate<br>
> Strict-Transport-Security: max-age=2592000<br>
> X-Frame-Options: deny<br>
> Set-Cookie: logged_in=no; domain=.github.com; path=/; expires=Tue, 01-Mar-2033 18:26:41 GMT<br>
> Set-Cookie: _gh_sess=BAh7Bz...d25af3; path=/; expires=Sun, 01-Jan-2023 00:00:00 GMT; secure; HttpOnly<br>
> X-Runtime: <b>6</b><br>
> ETag: "<b>de2ca9b0eba7131fdf5de1d6fcb461f4</b>"<br>
> Content-<b>Encoding: gzip</b><br>

And it highlights the new response header ``Content-Encoding`` ( instead of sending ``Content-Length`` ).

### TODO

- Add clear command
- Make way of highlighting configurable
- Find similar lines to provide a better highlighting
- Add option to hide unchanged lines

 

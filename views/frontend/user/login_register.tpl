<section class="header_text sub">
<img class="pageBanner" src="themes/images/pageBanner.png" alt="New products" >
  <h4><span>Login or Regsiter</span></h4>
</section>      
<section class="main-content">        
  <div class="row">
    <div class="span5">         
      <h4 class="title"><span class="text"><strong>Login</strong> Form</span></h4>
      {{.flash.error}}
      <form action="/login" method="post">
        {{.xsrfdata}}
        <fieldset>
          <div class="control-group">
            <label class="control-label">Username</label>
            <div class="controls">
              <input type="text" placeholder="Enter your username" id="username" name="username" class="input-xlarge">
            </div>
          </div>
          <div class="control-group">
            <label class="control-label">Username</label>
            <div class="controls">
              <input type="password" placeholder="Enter your password" id="password" name="password" class="input-xlarge">
            </div>
          </div>
          <div class="control-group">
            <input tabindex="3" class="btn btn-inverse large" type="submit" value="Sign into your account">
            <hr>
            <p class="reset">Recover your <a tabindex="4" href="#" title="Recover your username or password">username or password</a></p>
          </div>
        </fieldset>
      </form>       
    </div>
    {{if .Message}}
    <p class="text-danger">{{.Message}}</p>
    {{end}}
  
    <div class="span7">         
      <h4 class="title"><span class="text"><strong>Register</strong> Form</span></h4>
      <form action="/registerHandler" method="post" class="form-stacked">
        {{.xsrfdata}}
        <fieldset>
          <div class="control-group">
            {{ if .RegisterInfo.Errors.Name }}
            <p class="error">{{ .RegisterInfo.Errors.Name }}</p>
            {{ end }}
            <label class="control-label">Username</label>
            <div class="controls">
              <input type="text" name="userName" placeholder="Enter your username" class="input-xlarge">
            </div>
          </div>
          <div class="control-group">
            {{ if .RegisterInfo.Errors.Email }}
            <p class="error text-danger">{{ .RegisterInfo.Errors.Email }}</p>
            {{ end }}
            <label class="control-label">Email address:</label>
            <div class="controls">
              <input type="text" name="email" placeholder="Enter your email" class="input-xlarge">
            </div>
          </div>
          <div class="control-group">
            {{ if .RegisterInfo.Errors.Password }}
            <p class="error">{{ .RegisterInfo.Errors.Password }}</p>
            {{ end }}
            <label class="control-label">Password:</label>
            <div class="controls">
              <input type="password" name="password" placeholder="Enter your password" class="input-xlarge">
            </div>
          </div>  
          <div class="control-group">
            {{ if .RegisterInfo.Errors.Phone }}
            <p class="error">{{ .RegisterInfo.Errors.Phone }}</p>
            {{ end }}
            <label class="control-label">Phone Number:</label>
            <div class="controls">
              <input type="text" name="phoneNumber" placeholder="Enter your phone number" class="input-xlarge">
            </div>
          </div>
          <div class="control-group">
            {{ if .RegisterInfo.Errors.Address }}
            <p class="error">{{ .RegisterInfo.Errors.Address }}</p>
            {{ end }}
            <label class="control-label">Address:</label>
            <div class="controls">
              <input type="text" name="address" placeholder="Enter your address" class="input-xlarge">
            </div>
          </div>                                    
          <div class="control-group">
            <p>Now that we know who you are. I'm not a mistake! In a comic, you know how you can tell who the arch-villain's going to be?</p>
          </div>
          <hr>
          <div class="actions"><input tabindex="9" class="btn btn-inverse large" type="submit" value="Create your account"></div>
        </fieldset>
      </form>         
    </div>
    </div>    
</section>      

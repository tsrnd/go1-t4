<style>
  .user_profile{
    margin: 0 100px;
  }
</style>

<section class="main-content">
  <div class="user_profile">
  <h4 class="title"><span class="text">My <strong>Account</strong> Profile</span></h4>
  <div class="row">
      <div class="span3">
      <div class="col-sm-8">
          <ul class="list-group">
              <a href="/userProfile"><img src="/static/css/themes/images/avatar1.png" class="site_avatr" alt=""></a>
              <li class="list-group-item text-right">
                <span class="pull-left"><strong class="">UserName:</strong></span> {{ .User.UserName }}
              </li>
              <li class="list-group-item text-right">
                <span class="pull-left"><strong class="">Email:</strong></span><a href="#"> {{ .User.Email }}</a>
              </li>
              <li class="list-group-item text-right">
                <span class="pull-left"><strong class="">Gender:</strong></span> {{ .User.Gender }}
              </li>
              <li class="list-group-item text-right">
                <span class="pull-left"><strong class="">Phone Number:</strong></span> {{ .User.Phone }}
              </li>
              <li class="list-group-item text-right">
                <span class="pull-left"><strong class="">Address:</strong></span> {{ .User.Address }}
              </li>
              <li class="list-group-item text-right">
                <span class="pull-left"><strong class="">Avatar:</strong></span> {{ .User.Avatar }}
              </li>
          </ul>
        </div>
      </div>
    </div>
</div>
</section> 

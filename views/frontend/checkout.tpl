<section class="header_text sub">
<img class="pageBanner" src="/static/css/themes/images/pageBanner.png" alt="New products" >
  <h4><span>Check Out</span></h4>
</section>  
<section class="main-content">
  <div class="row">
    <div class="span12">
        <h4 id="js-message" class="text-center">{{ .Message }}</h4>                    
      <div class="accordion" id="accordion2">
        <div class="accordion-group">
          <div class="accordion-heading">
            <a class="accordion-toggle" data-toggle="collapse" data-parent="#accordion2" href="#collapseOne">Checkout Options</a>
          </div>
          <div id="collapseOne" class="accordion-body in collapse">
            <div class="accordion-inner">
              <div class="row-fluid">
                <div class="span3"></div>
                <div class="span6">
                <h4>Returning Customer</h4>
                <p>Customer infomation </p>
                  <form method="post" action="/order" id="js-form-sbm">
                    {{.xsrfdata}}
                    <fieldset>
                      <div class="control-group">
                        <label class="control-label"><strong>Username</strong></label>
                        <div class="controls">
                          <input type="text" value="{{ .Name }}" readonly="readonly" class="input-xlarge">
                        </div>
                      </div>
                      <div id="js-list-product"></div>
                      <div class="control-group">
                        <label class="control-label"><strong>TotalMoney</strong></label>
                        <div class="controls">
                          <input type="text" name="total_money" id="total_money" readonly="readonly" class="input-xlarge">
                        </div>
                      </div>
                      <div class="control-group">
                        <label class="control-label"><strong>Payment method</strong></label>
                        <div class="controls">
                          <select name="payment_id" id="">
                            {{ range $payment := .Payments }}
                            <option value="{{ $payment.ID }}">{{ $payment.Method }}</option>
                            {{ end }}
                          </select>
                        </div>
                      </div>
                      
                      <button type="submit" id="checkout" class="btn btn-inverse">Checkout</p>
                    </fieldset>
                  </form>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>        
    </div>
  </div>
  <div class="modal fade cls-modal-booking" id="success-modal" role="dialog">
    <div class="modal-dialog">
    
      <!-- Modal content-->
      <div class="modal-content">
        <div class="modal-body">
          <div class="span1">
            <img class="" style="width:110px" src="/static/css/themes/images/check.svg" alt="">
          </div>
          <div class="">
            <p class="cls-modal-booking-text text-center text-primary">Order Success</p>
            <p class="text-center text-primary">Thank you for trusting us to be your source for top quality sporting. </p>
          </div>
        </div>
        <div class="modal-footer cls-modal-booking-footer">
          <button id="js-btn-close" class="btn btn-default">Yes, I known</button>
        </div>
      </div>
      
    </div>
  </div>
</section>

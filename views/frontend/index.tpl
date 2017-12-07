<section  class="homepage-slider" id="home-slider">
  <div class="flexslider">
    <ul class="slides">
      <li>
        <img src="/static/css/themes/images/carousel/banner-1.jpg" alt="" />
      </li>
      <li>
        <img src="/static/css/themes/images/carousel/banner-2.jpg" alt="" />
        <div class="intro">
          <h1>Mid season sale</h1>
          <p><span>Up to 50% Off</span></p>
          <p><span>On selected items online and in stores</span></p>
        </div>
      </li>
    </ul>
  </div>      
</section>
<section class="header_text">
  We stand for top quality templates. Our genuine developers always optimized bootstrap commercial templates. 
  <br/>Don't miss to use our cheap abd best bootstrap templates.
</section>
<section class="main-content">
  <div class="row">
    <div class="span12">                          
      <div class="row">
        <div class="span12">
          <h4 class="title">
            <span class="pull-left"><span class="text"><span class="line">Latest <strong>Products</strong></span></span></span>
          </h4>
          <div id="myCarousel" class="myCarousel carousel slide">
            <div class="carousel-inner">
              <div class="active item">
                  <ul class="thumbnails listing-products">
                      {{ range $index, $product := .LatesProducts }}
                        <li class="span3">
                          <div class="product-box">
                            <span class="sale_tag"></span>                        
                            <a href="/detail/{{ $product.ID }}"><img alt="" src="/{{ if $product.Images }}{{ (index $product.Images 0).URL }}{{ else }}static/css/themes/images/ladies/6.jpg {{ end }}"></a><br/>
                            <a href="/detail/{{ $product.ID }}" class="title">{{ $product.Name }}</a><br/>
                            <a href="#" class="category">SIZE: {{ $product.Size }}</a>
                            <p class="price">{{ $product.Price }}$</p>
                          </div>
                        </li>
                      {{ end }}
                    </ul>
              </div>
            </div>              
          </div>
        </div>            
      </div>
      <br/>
      <div class="row">
        <div class="span12">
          <h4 class="title">
            <span class="pull-left"><span class="text"><span class="line">BEST <strong>SELLER</strong></span></span></span>
            <!-- <span class="pull-right">
              <a class="left button" href="#myCarousel-2" data-slide="prev"></a><a class="right button" href="#myCarousel-2" data-slide="next"></a>
            </span> -->
          </h4>
          <div id="myCarousel-2" class="myCarousel carousel slide">
            <div class="carousel-inner">
              <div class="active item">
                  <ul class="thumbnails listing-products">
                      {{ range $index, $product := .BestSeller }}
                        <li class="span3">
                          <div class="product-box">
                            <span class="sale_tag"></span>                        
                            <a href="/detail/{{ $product.ID }}"><img alt="" src="/{{ if $product.Images }}{{ (index $product.Images 0).URL }}{{ else }}static/css/themes/images/ladies/6.jpg {{ end }}"></a><br/>
                            <a href="/detail/{{ $product.ID }}" class="title">{{ $product.Name }}</a><br/>
                            <a href="#" class="category">SIZE: {{ $product.Size }}</a>
                            <p class="price">{{ $product.Price }}$</p>
                          </div>
                        </li>
                      {{ end }}
                    </ul>
              </div>
              <div class="item">
              </div>
            </div>              
          </div>
        </div>            
      </div>
      <div class="row feature_box">           
        <div class="span4">
          <div class="service">
            <div class="responsive">  
              <img src="/static/css/themes/images/feature_img_2.png" alt="" />
              <h4>MODERN <strong>DESIGN</strong></h4>
              <p>Lorem Ipsum is simply dummy text of the printing and printing industry unknown printer.</p>                  
            </div>
          </div>
        </div>
        <div class="span4"> 
          <div class="service">
            <div class="customize">     
              <img src="/static/css/themes/images/feature_img_1.png" alt="" />
              <h4>FREE <strong>SHIPPING</strong></h4>
              <p>Lorem Ipsum is simply dummy text of the printing and printing industry unknown printer.</p>
            </div>
          </div>
        </div>
        <div class="span4">
          <div class="service">
            <div class="support"> 
              <img src="/static/css/themes/images/feature_img_3.png" alt="" />
              <h4>24/7 LIVE <strong>SUPPORT</strong></h4>
              <p>Lorem Ipsum is simply dummy text of the printing and printing industry unknown printer.</p>
            </div>
          </div>
        </div>  
      </div>    
    </div>        
  </div>
</section>
<section class="our_client">
  <h4 class="title"><span class="text">Manufactures</span></h4>
  <div class="row">         
    <div class="span2">
      <a href="#"><img alt="" src="/static/css/themes/images/clients/14.png"></a>
    </div>
    <div class="span2">
      <a href="#"><img alt="" src="/static/css/themes/images/clients/35.png"></a>
    </div>
    <div class="span2">
      <a href="#"><img alt="" src="/static/css/themes/images/clients/1.png"></a>
    </div>
    <div class="span2">
      <a href="#"><img alt="" src="/static/css/themes/images/clients/2.png"></a>
    </div>
    <div class="span2">
      <a href="#"><img alt="" src="/static/css/themes/images/clients/3.png"></a>
    </div>
    <div class="span2">
      <a href="#"><img alt="" src="/static/css/themes/images/clients/4.png"></a>
    </div>
  </div>
</section>

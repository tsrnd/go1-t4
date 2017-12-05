$(document).ready(function() {
  hiddenButtonAdd()
  showNumberInCart()
  showCart()
  showCheckout()
  $('#js-btn-add-cart').on('click', function(){
    addCart()
  });
});

function addCart(){
      var OldCart = new Array
      var product = new Object()

      product.Name = $('#js-pr-name').html()
      product.Id = $('#js-pr-id').html()
      product.Price = $('#js-pr-price').html()
      product.Image = $('#js-pr-image').attr('src')
      product.Quantity = $('#js-pr-quantity').val()

      if (getCart() != null){
        OldCart = getCart()
        OldCart.forEach(function(item){
          if (item.Id == product.Id) {
            product.Quantity = parseInt(product.Quantity) + parseInt(item.Quantity)
            OldCart.splice($.inArray(item, OldCart),1);
          }
        })
      }
  
      OldCart.push(product)
      var s = JSON.stringify(OldCart);
  
      localStorage.setItem('cart', s);
      $('#success-modal').modal('show')
      $('#js-btn-close').on('click', function(){
        $('#success-modal').modal('hide')
        showNumberInCart()
      })
}

function getCart(){
  var OldCart = new Array
  OldCart = JSON.parse(localStorage.getItem("cart"))
  return OldCart
}

function hiddenButtonAdd(){
  if ($('#js-pr-quantity').val() < 1) {
    $('#js-btn-add-cart').attr("disabled","disabled")
  }
  $('#js-pr-quantity').on('input', function(){
    if ($('#js-pr-quantity').val() < 1) {
      $('#js-btn-add-cart').attr("disabled","disabled")
    } else {
      $('#js-btn-add-cart').removeAttr('disabled')
    }
  });
}

function showNumberInCart(){
  var Cart = getCart()
  if(Cart != null){
    $('#js-cart-count').html(Cart.length)    
  }
}

function showCart(){
  var rs = ""
  var form = $('#js-form-sbm')
  var products = getCart()
  if (products == null){
    $('#js-show-cart').html('<h4>Your cart is empty<h4>')
  } else {
    products.forEach(function(item){
      rs += '<tr>'
      rs += '<td><a href="/detail/'+ item.Id +'"><img alt="" src="'+ item.Image +'"></a></td>'
      rs += '<td> '+ item.Name+ '</td>'
      rs += '<td><p>'+ item.Quantity +'</p></td>'
      rs += '<td>$'+ item.Price +'</td>'
      rs += '<td>$'+ item.Price*item.Quantity +'</td>'
      rs += '<td><button class="btn btn-danger pull-right" onClick="removeItemInCart('+ item.Id +')">X</button></td>'
      rs += '</tr>'
    })
  }
  $('#js-show-cart').html(rs)
}

function removeItemInCart(id){
  OldCart = getCart()
  OldCart.forEach(function(item){
    if (getCart() != null){
      OldCart = getCart()
      OldCart.forEach(function(item){
        if (item.Id == id) {
          OldCart.splice($.inArray(item, OldCart),1);
        }
      })
    }
    var s = JSON.stringify(OldCart);
    localStorage.setItem('cart', s);
    showNumberInCart()
    showCart()
    })
}

function hiddenButtonAdd(){
  if ($('#js-pr-quantity').val() < 1) {
    $('#js-btn-add-cart').attr("disabled","disabled")
  }
  $('#js-pr-quantity').on('input', function(){
    if ($('#js-pr-quantity').val() < 1) {
      $('#js-btn-add-cart').attr("disabled","disabled")
    } else {
      $('#js-btn-add-cart').removeAttr('disabled')
    }
    console.log($('#js-pr-quantity').val())
  });
}

function showNumberInCart(){
  var Cart = getCart()
  if(Cart != null){
    $('#js-cart-count').html(Cart.length)    
  }
}

function showCheckout(){
  var rs = ""
  var Total = 0
  var products = getCart()
  if (products == null){
    $('#js-list-product').html('<h4>Your cart is empty<h4>')
  } else {
    products.forEach(function(item){
  
      var p = document.createElement('p')
      p.append('Product: '+item.Name+ '---Quantity: ' + item.Quantity + '---Total Money: '+item.Quantity*item.Price)
      Total += item.Quantity*item.Price

      // create form
      var e1 = document.createElement('input')
      var e2 = document.createElement('input')
      e1.name = 'product_id'; e1.hidden = 'hidden'
      e2.name = 'quantity'; e2.hidden = 'hidden'
      e1.value = item.Id
      e2.value = item.Quantity
      $('#js-list-product').append(e1)
      $('#js-list-product').append(e2)
      $('#js-list-product').append(p)
    })
    $('#total_money').val(Total)
  }
}

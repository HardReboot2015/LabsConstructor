'use strict';

let canvas=document.getElementById("canvas");
let ctx=canvas.getContext("2d");
let BB=canvas.getBoundingClientRect();
let offsetX=BB.left;
let offsetY=BB.top;
let WIDTH = canvas.width;
let HEIGHT = canvas.height;




// drag related letiables
let dragok = false;
let startX;
let startY;
let currentFigure;

let elements_img = [];
//получаем єлементы от страницы
fetch('/GetElements' ).then((responce) => {
  return responce.json();
}).then((data)=>{
  for (let i = 0; i < data.length; i++){
    elements_img[i] = data[i];

  }
})
//получаем кнопки со страницы и даем им функцию онклик

let button_images = document.querySelectorAll(".image_menu");
Array.from(button_images).forEach(function (button_image){
  button_image.addEventListener('click', function (e) {
    // alert(e.target.id);
    let img = document.getElementById(e.target.id);
    addimage(e.target.id, img);
  })
})
let images = [];
let imgId = 0;


canvas.onmousedown = myDown;
canvas.onmouseup = myUp;
// canvas.onmousemove = myMove;
canvas.ondblclick = myDblclick;
// canvas.onclick = myClick();
// image-png.onclick = imageOnclick();

//
// //рисуем линию
// function drawLine(){
//   ctx.moveTo(offsetX, offsetY);
//   ctx.lineTo(offsetX, offsetY)
// }



function addimage(id, img){
  if (images.length === 0){
  images.push({id:1,x:20,y:20, width:111, height:65, rotate: 0,img:img, image_data: elements_img[id]});
  }else{
    images.push({id:images.length+1,x:20, y:20,width:111, height:65, rotate :0,img:img,image_data: elements_img[id]});
  }
  // console.log(images[0].image.Svg)
  draw();
}

function clear() {
  ctx.clearRect(0, 0, WIDTH, HEIGHT);
}
function clearElem(x, y, width, height){
  ctx.clearRect(x, y, width, height);
  ctx.rect(x, y, width, height);
  ctx.stroke();
}

function drawImg(i){
  //
  // var mx=parseInt(e.clientX-offsetX);
  // var my=parseInt(e.clientX-offsetY);


}

function rotateImg(i) {
  
  ctx.save();
  
  let x = i.x + i.width/2;
  let y = i.x + i.height/2;
  
  if(i.rotate === 0){
    clearElem(i.x, i.y, i.width, i.height);
  }else{
    clearElem(i.x, i.y, i.height, i.width);
  }
  ctx.translate(i.x, i.y)
  
  if (i.rotate === 0){

    i.rotate = 1;
  }else{
    ctx.rotate(90 * Math.PI/180);
    i.rotate = 0;
  }

  ctx.drawImage(i.img, -(i.width/2), -(i.height/2), i.width, i.height);
  
  ctx.restore();

}

// function rotateDegree() {
  //   clear();
  //   for(let i=0; i<images.length;i++) {
    //     rotateImg(images[i]);
    //   }
    // }


    
    function draw() {
      clear();
      for(let i=0; i < images.length;i++) {
        // if (images[i].rotate === 1){
          rotateImg(images[i]);
        // }else{
        //   drawImg(images[i]);
        // }

      }
    }


line.onclick = () =>{
  currentFigure ='line';
}
// function myClick(){
//   if (currentFigure === 'line'){
//
//   }
// }
    // handle mousedown events
    function myDown(e){
      
      // tell the browser we're handling this mouse event
      e.preventDefault();
      e.stopPropagation();
      
      // get the current mouse position
      var mx=parseInt(e.clientX-offsetX);
      var my=parseInt(e.clientY-offsetY);
      
      // test each shape to see if mouse is inside
      dragok=false;
      for(var i=0;i<images.length;i++){
        var m=images[i];
        // test if the mouse is inside this rect
       
          if(mx>m.x && mx<m.x+m.width && my>m.y && my<m.y+m.height){
          // if yes, set that rects isDragging=true
          dragok=true;
          m.isDragging=true;
        
      }
      }
      startX=mx;
      startY=my;
    }
    function myDblclick(e){
      
      // tell the browser we're handling this mouse event
      e.preventDefault();
      e.stopPropagation();
      
      // get the current mouse position
      var mx=parseInt(e.clientX-offsetX);
      var my=parseInt(e.clientY-offsetY);
      
      // test each shape to see if mouse is inside
      dragok=false;
      for(var i=0;i<images.length;i++){
        var m=images[i];
        
        // test if the mouse is inside this rect
        if (m.rotateparam==0){
        if(mx>m.x && mx<m.x+m.width && my>m.y && my<m.y+m.height){
          // if yes, set that rects isDragging=true
          // dragok=true;
          // m.isDragging=true;
          rotateImg(m);
        }
      }else{
          if (mx>m.x && mx<m.x+m.height && my>m.y && my<m.y+m.width){
            rotateImg(m);

            
          }
        
      }
      }
      startX=mx;
      startY=my;
    }
// handle mouseup events
function myUp(e){
  // tell the browser we're handling this mouse event
  e.preventDefault();
  e.stopPropagation();
  
  // clear all the dragging flags
  dragok = false;
  for(var i=0;i<images.length;i++){
    images[i].isDragging=false;
  }
}


// handle mouse moves
function myMove(e){
  // if we're dragging anything...
  if (dragok){
    
    // tell the browser we're handling this mouse event
    e.preventDefault();
    e.stopPropagation();
    
    // get the current mouse position
    var mx=parseInt(e.clientX-offsetX);
    var my=parseInt(e.clientY-offsetY);
    
    // calculate the distance the mouse has moved
    // since the last mousemove
    var dx=mx-startX;
    var dy=my-startY;
    
    // move each rect that isDragging 
    // by the distance the mouse has moved
    // since the last mousemove
    for(var i=0;i<images.length;i++){
      var m=images[i];
      if(m.isDragging){
        m.x+=dx;
        m.y+=dy;
      }
    }
    
    // redraw the scene with the new rect positions
    draw();
    
    // reset the starting mouse position for the next mousemove
    startX=mx;
    startY=my;
    
  }
}

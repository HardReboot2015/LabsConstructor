"use strict";

function arr() {
    let isDone = confirm('Do u want to end trial?');
    if(isDone === true) {
        window.location.href = "/cabinet/user";
    }else return;
}

end.onclick = arr;

//
// let img = document.getElementById("line");
//
// let context = canvas.getContext('2d');
// let cnvs = document.getElementById("canvas");
//
//
// let add = () => context.drawImage(img, 70, 20, 111, 65);
//
// context.lineWidth = 1; // толщина линии
//
// context.moveTo(0,0)
//
// cnvs.onclick = function () {
//    let x=event.offsetX;
//    let y=event.offsetY;
// context.lineTo(x, y); //рисуем линию
// context.stroke();
// }

//

/*let photo = document.getElementById('photo');

photo.onmousedown = function(e) { // 1. отследить нажатие

  // подготовить к перемещению
  // 2. разместить на том же месте, но в абсолютных координатах
  photo.style.position = 'absolute';
  moveAt(e);
  // переместим в body, чтобы мяч был точно не внутри position:relative
  document.body.appendChild(photo);

  photo.style.zIndex = 1000; // показывать мяч над другими элементами

  // передвинуть мяч под координаты курсора
  // и сдвинуть на половину ширины/высоты для центрирования
  function moveAt(e) {
    photo.style.left = e.pageX - photo.offsetWidth / 2 + 'px';
    photo.style.top = e.pageY - photo.offsetHeight / 2 + 'px';
  }

  // 3, перемещать по экрану
  document.onmousemove = function(e) {
    moveAt(e);
  }

  // 4. отследить окончание переноса
  photo.onmouseup = function() {
    document.onmousemove = null;
    photo.onmouseup = null;
  }
}
*/

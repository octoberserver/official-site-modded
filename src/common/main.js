{
    let logoImage = document.getElementsByClassName("nav-logo-img")[0]
    let a = logoImage.parentElement;
    a.addEventListener("mouseover", e => logoImage.src = "common/logo-blink-2.gif");
    a.addEventListener("mouseout", e => logoImage.src = "common/logo.png");
}

{
    let logoImage = document.querySelector("[nav-logo-img]");
    let a = logoImage.parentElement;
    a.addEventListener("mouseover", e => logoImage.src = "common/logo-blink-2.gif");
    a.addEventListener("mouseout", e => logoImage.src = "common/logo.png");
}

{
    let logoImage = document.querySelector("[nav-logo-img]");
    let a = logoImage.parentElement;
    a.addEventListener("mouseover", e => logoImage.src = "common/images/icons/logo-blink-2.gif");
    a.addEventListener("mouseout", e => logoImage.src = "common/images/icons/logo.png");
}

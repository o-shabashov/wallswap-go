/*Theme    : Quick
 * Author  : Design_mylife
 * Version : V1.0
 *
 */

/* ==============================================
 Sticky Navbar
 =============================================== */

$(document).ready(function() {
    $(".navbar-sticky").sticky({topSpacing: 0});
});



/* ==============================================
 main flex slider
 =============================================== */
$(document).ready(function() {
    $('.main-flex-slider').flexslider({
        slideshowSpeed: 5000,
        directionNav: false,
        animation: "fade",
        controlNav: false
    });
});


/* ==============================================
 Auto Close Responsive Navbar on Click
 =============================================== */

function close_toggle() {
    if ($(window).width() <= 768) {
        $('.navbar-collapse a').on('click', function() {
            $('.navbar-collapse').collapse('hide');
        });
    }
    else {
        $('.navbar .navbar-default a').off('click');
    }
}
close_toggle();

$(window).resize(close_toggle);


/* ===================================================================
 TWEETIE -  TWITTER FEED PLUGIN THAT WORKS WITH NEW Twitter 1.1 API
 ==================================================================== */
$('.tweet').twittie({
    apiPath: '/twit-api/tweet.php',
    dateFormat: '%b. %d, %Y',
    template: '{{tweet}} <div class="date">{{date}}</div> <a href="{{url}}"{{screen_name}}',
    count: 2
});

/***================================================== */
$('.chart').each(function () {
    var $this = $(this);
    var color = $(this).data('scale-color');

    setTimeout(function () {
        $this.filter(':visible').waypoint(function (direction) {
            $(this).easyPieChart({
                barColor: color,
                trackColor: '#fff',
                onStep: function (from, to, percent) {
                    jQuery(this.el).find('.percent').text(Math.round(percent));
                }
            });
        }, {offset: '100%'});
    }, 500);

});

//owl carousel for testimonials
$(document).ready(function() {

    $("#testi-carousel,#work-slide").owlCarousel({
        // Most important owl features
        items: 1,
        itemsCustom: false,
        itemsDesktop: [1199, 1],
        itemsDesktopSmall: [980, 1],
        itemsTablet: [768, 1],
        itemsTabletSmall: false,
        itemsMobile: [479, 1],
        singleItem: false,
        startDragging: true,
        autoPlay: true
    });

});

/*=========================*/
/*========portfolio mix====*/
/*==========================*/
$('#grid').mixitup();


/*=========================*/
/*========on hover dropdown navigation====*/
/*==========================*/


/************parallax*********************/
$(function() {
    $.stellar({
        horizontalScrolling: false
    });
});


/* ==============================================
 Counter Up
 =============================================== */
jQuery(document).ready(function($) {
    $('.counter').counterUp({
        delay: 100,
        time: 800
    });
});

/* ==============================================
 WOW plugin triggers animate.css on scroll
 =============================================== */

var wow = new WOW(
    {
        boxClass: 'wow', // animated element css class (default is wow)
        animateClass: 'animated', // animation css class (default is animated)
        offset: 100, // distance to the element when triggering the animation (default is 0)
        mobile: false        // trigger animations on mobile devices (true is default)
    }
);
wow.init();


//MAGNIFIC POPUP
$('.show-image').magnificPopup({type: 'image'});


//smooth scroll
$(function() {
    $('.scrollto a[href*=#]:not([href=#])').click(function() {
        if (location.pathname.replace(/^\//,'') == this.pathname.replace(/^\//,'') && location.hostname == this.hostname) {

            var target = $(this.hash);
            target = target.length ? target : $('[name=' + this.hash.slice(1) +']');
            if (target.length) {
                $('html,body').animate({
                    scrollTop: target.offset().top - 50
                }, 1000);
                return false;
            }
        }
    });
});

// Backstretch - Slider on Background
$(".fullscreen").backstretch([
    "img/showcase-5.jpg",
    "img/showcase-2.jpg",
    "img/showcase-3.jpg"
], {duration: 5000, fade: 1000});

//back to top
$(document).ready(function(){

    //Check to see if the window is top if not then display button
    $(window).scroll(function(){
        if ($(this).scrollTop() > 800) {
            $('.scrollToTop').fadeIn();
        } else {
            $('.scrollToTop').fadeOut();
        }
    });

    //Click event to scroll to top
    $('.scrollToTop').click(function(){
        $('html, body').animate({scrollTop : 0},800);
        return false;
    });

});


(function ($) {
    $(document).ready(function () {
        var body = $('body');

        // Add tag to wallpaper on change event
        body.on('change', 'input.tagsinput', function (event) {
            var _this = this;
            var wallpaperId = ($(this).attr('data-wallpaper-id'));
            var url = ($(this).attr('data-tag-url'));

            $.ajax({
                type: 'POST',
                url : url + '/add',
                data: {tag: $(_this).val(), wallpaperId: wallpaperId},

                success: function (data) {
                    $('.tags-list').append(data);
                }
            });
            event.preventDefault();
        });

        // Add tag to wallpaper on typeahead:selected event
        body.on('typeahead:selected', 'input.tagsinput', function (object, datum) {
            var _this = this;
            var wallpaperId = ($(this).attr('data-wallpaper-id'));
            var url = ($(this).attr('data-tag-url'));

            $.ajax({
                type: 'POST',
                url : url + '/add',
                data: {tag: $(_this).val(), wallpaperId: wallpaperId},

                success: function (data) {
                    $('.tags-list').append(data);
                }
            });
        });

        // Wallpaper set purity
        $('.set-purity').click(function () {
            var wallpaperId = ($(this).attr('data-wallpaper-id'));
            var purity = ($(this).attr('data-purity'));
            var url = ($(this).attr('data-url'));
            var button = $(this);
            var img = $('.img-responsive');
            $.ajax({
                type   : "POST",
                url    : url,
                data   : {wallpaperId: wallpaperId, purity: purity},
                success: function (data) {
                    img.removeClass('img-sfw');
                    img.removeClass('img-sketchy');
                    img.removeClass('img-nsfw');
                    img.addClass('img-' + purity);

                    $('.set-purity').removeClass('active');

                    button.addClass('active');
                }
            });
        });

        // All wallpapers in queue set purity
        $('.all-set-purity').click(function () {
            var purity = ($(this).attr('data-purity'));
            var url = ($(this).attr('data-url'));
            var button = $(this);
            var img = $('.img-responsive');
            $.ajax({
                type   : "GET",
                url    : url,
                data   : {purity: purity},
                success: function (data) {
                    img.removeClass('img-sfw');
                    img.removeClass('img-sketchy');
                    img.removeClass('img-nsfw');
                    img.addClass('img-' + purity);

                    $('.all-set-purity').removeClass('active');

                    button.addClass('active');
                }
            });
        });

        // Tag set purity
        $('.tag-set-purity').click(function () {
            var tagId = ($(this).attr('data-tag-id'));
            var purity = ($(this).attr('data-purity'));
            var url = ($(this).attr('data-url'));
            var button = $(this);
            $.ajax({
                type   : "POST",
                url    : url,
                data   : {tagId: tagId, purity: purity},
                success: function (data) {
                    $('.set-purity').removeClass('active');

                    button.addClass('active');
                }
            });
        });

        // Tag input
        var tagsInputSelector = $('.tagsinput');
        if (tagsInputSelector.length != 0) {
            var tags = new Bloodhound({
                datumTokenizer: Bloodhound.tokenizers.obj.whitespace('name'),
                queryTokenizer: Bloodhound.tokenizers.whitespace,
                prefetch      : {
                    url   : '/js/tags.json',
                    filter: function (list) {
                        return $.map(list, function (skill) {
                            return {name: skill};
                        });
                    }
                }
            });
            tags.initialize();

            tagsInputSelector.typeahead(null, {
                name      : 'tags',
                displayKey: 'name',
                valueKey  : 'name',
                source    : tags.ttAdapter()
            });

        }

        // Lazy load wallpapers
        function lazyLoadFunc() {
            $(document).on("ready ajaxComplete", function () {
                $("img.lazy").show().lazyload({
                    effect      : "fadeIn",
                    effect_speed: 500,
                    threshold   : 50,
                    placeholder : "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAEElEQVR42gEFAPr/AAoKCv8BXgEeaJsaowAAAABJRU5ErkJggg=="
                }).removeClass("lazy");
            });
        }
        lazyLoadFunc();

        // Ajax infinite scroll
        $('.list-view').jscroll({
            padding       : 50,
            loadingHtml   : '<div class="row"><div class="col-lg-12 col-md-12 col-sm-12 col-xs-12 text-center"><div class="loader">Loading...</div></div></div>',
            nextSelector  : 'ul.pagination li.active + li a:last',
            pagingSelector: '.pagination',
            callback      : lazyLoadFunc()
        });

        // Bootstrap switch
        if ($("input[type='checkbox']").length) {
            $("input[type='checkbox']").bootstrapSwitch();
        }
    });
}(window.jQuery));
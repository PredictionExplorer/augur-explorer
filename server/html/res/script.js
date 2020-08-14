(function(){
    const initPreventBehavior = () => {
        $(document).on('click', 'a[href="#"]', function(ev) {
            ev.preventDefault();
        });
    };


    const showMoreDetails = () => {
        $('[deatils-more-js]').on('click', (ev) => {
            if($(ev.currentTarget).hasClass('is-active')) {
                $(ev.currentTarget).removeClass('is-active');
                $('[deatils-desc-js]').removeClass('is-show').css({
                    height: 40
                });
            } else {
                $(ev.currentTarget).addClass('is-active');
                $('[deatils-desc-js]').addClass('is-show').css({
                    height: $('.details__description').outerHeight(true)
                });
            }
        });
    };

    const detailsCollapse = () => {
        $('[details-collapse-js]').on('click', (ev) => {
            $(ev.currentTarget).toggleClass('is-active');
            $(ev.currentTarget).siblings('[details-collapse-body-js]').slideToggle(300);
        });
    };

    const detailsStatAdvanced = () => {
        $('[stat-head-js]').on('click', (ev) => {
            $(ev.currentTarget).toggleClass('is-active');
            $(ev.currentTarget).siblings('[stat-body-js]').slideToggle(300);
        });
    };

    const loadMoreOrders = () => {
        $('[load-more-js]').on('click', (ev) => {
            $(ev.currentTarget).fadeOut(300);
            $('.details__orders table tbody tr:hidden').fadeIn(500);
        });
    };


    window.addEventListener('load', (ev) => {
        initPreventBehavior();
        showMoreDetails();
        detailsCollapse();
        detailsStatAdvanced();
        loadMoreOrders();
    }, false);
})();
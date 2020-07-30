function getSingleCardData(data, renderFunction) {
    console.log('Start Rendering')
    for(let i = 0; i <= data.MarketIDs.length; i++  ) {
        Ajax_GET(`/api/mkt_card/${data.MarketIDs[i]}`, renderFunction);
    }
}
function renderCardIndex(data) {
    console.log('Draw')
    data = JSON.parse(data);
    console.log(data)
    document.querySelector('#ajax-content').innerHTML += `
    <div class="col-sm-6 col-md-4">
        <div class="card">
            <div class="card__title">
                <a href="#">${data.MarketInfo.Description}</a>
                <div class="card__title-down">
                    <span>Ends in 2 days</span><span class="tag">${data.MarketInfo.MktTypeStr}</span>
                </div>
            </div>
            <div class="card__info">
                <div class="card__all-left">
                    <div class="card__percents">
                        <span class="answer">Yes</span>
                        <span class="answer">No</span>
                    </div>
                    <div class="card__number">
                        <span class="percent-yes">100%</span>
                        <span class="percent-no">0%</span>
                    </div>
                </div>
                <div class="card__all-right">
                    <div class="card__volume">
                        <span>Total volume</span>
                        <span>Money at stake</span>
                    </div>
                    <div class="card__money">
                        <span>${data.MarketInfo.CurVolume}</span>
                        <span>224.5</span>
                    </div>
                </div>
            </div>
            <p class="card__footer">And <a href="#">1 more</a> possible outcome</p>
        </div>
    </div>`
}
function renderCard(data) {
    console.log('Draw')
    data = JSON.parse(data);
    console.log(data)
    document.querySelector('#ajax-content').innerHTML += `

   <div class="col-md-6">
                <div class="market__info-box">
                    <div class="row">
                        <div class="col-sm-6"><p>Outcome</p></div>
                        <div class="col-sm-6"><p>Last Price</p></div>
                    </div>
                    <section id="multiple" data-accordion-group="">
                        <section data-accordion="" class="accordion-box">
                            <div data-control="">
                                <div class="row">
                                    <div class="col-sm-6">
                                        <div class="progress-title">
                                            <p>Yes</p>
                                            <p class= "progress-title-yes">90%</p>
                                        </div>
                                        <div class="w3-light-grey">
                                            <div class="w3-color-yes" style="width:90%"></div>
                                        </div>
                                    </div>
                                    <div class="col-sm-3">
                                        <p>$31</p>
                                    </div>
                                    <div class="col-sm-3">
                                       <span class="market__info-detail">View Details</span>
                                    </div>
                                </div>
                            </div>
                            <div data-content="" style="max-height: 60px; overflow: hidden;">
                                <div class="row">
                                    <div class="col-sm-3">
                                        <p>Bid price</p>
                                        <span>0.55</span>
                                    </div>
                                    <div class="col-sm-3">
                                        <p>Ask price</p>
                                        <span>1</span>
                                    </div>
                                    <div class="col-sm-3">
                                        <p>Volume</p>
                                        <span>${data.MarketInfo.CurVolume}</span>
                                    </div>
                                    <div class="col-sm-3">
                                        <a href="#">${data.MarketInfo.Description}</a>
                                    </div>
                                </div>
                            </div>
                        </section>

                        <section data-accordion="" class="accordion-box">
                            <div data-control="">
                                <div class="row">
                                    <div class="col-sm-6">
                                        <div class="progress-title">
                                            <p>No</p>
                                            <p class= "progress-title-no">10%</p>
                                        </div>
                                        <div class="w3-light-grey">
                                            <div class="w3-color-no" style="width:10%"></div>
                                        </div>
                                    </div>
                                    <div class="col-sm-3">
                                        <p>$68</p>
                                    </div>
                                    <div class="col-sm-3">
                                        <span class="market__info-detail">View Details</span>
                                    </div>
                                </div>
                            </div>
                            <div data-content="" style="max-height: 60px; overflow: hidden;">
                                <div class="row">
                                    <div class="col-sm-3">
                                        <p>Bid price</p>
                                        <span>0.55</span>
                                    </div>
                                    <div class="col-sm-3">
                                        <p>Ask price</p>
                                        <span>1</span>
                                    </div>
                                    <div class="col-sm-3">
                                        <p>Volume</p>
                                        <span>401</span>
                                    </div>
                                    <div class="col-sm-3">
                                        <a href="#">Depth chart</a>
                                    </div>
                                </div>
                            </div>
                        </section>

                        <section data-accordion="" class="accordion-box">
                            <div data-control="">
                                <div class="row">
                                    <div class="col-sm-6"><p>Invalid</p></div>
                                    <div class="col-sm-3">
                                        <p>$68</p>
                                    </div>
                                    <div class="col-sm-3">
                                        <span class="market__info-detail">View Details</span>
                                    </div>
                                </div>
                            </div>
                            <div data-content="" style="max-height: 60px; overflow: hidden;">
                                <div class="row">
                                    <div class="col-sm-3">
                                        <p>Bid price</p>
                                        <span>0.55</span>
                                    </div>
                                    <div class="col-sm-3">
                                        <p>Ask price</p>
                                        <span>1</span>
                                    </div>
                                    <div class="col-sm-3">
                                        <p>Volume</p>
                                        <span>401</span>
                                    </div>
                                    <div class="col-sm-3">
                                        <a href="#">Depth chart</a>
                                    </div>
                                </div>
                            </div>
                        </section>
                    </section>
                </div>
            </div>
   `
}
function Ajax_GET(url, success) {
    var xhr = window.XMLHttpRequest ? new XMLHttpRequest() : new ActiveXObject('Microsoft.XMLHTTP');
    xhr.open('GET', url);
    xhr.onreadystatechange = function() {
        if (xhr.readyState>3 && xhr.status==200) success(xhr.responseText);
    };
    xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest');
    xhr.send();
    return xhr;
}
function Ajax_GET_Markets(url, renderFunction) {
    var xhr = window.XMLHttpRequest ? new XMLHttpRequest() : new ActiveXObject('Microsoft.XMLHTTP');
    xhr.open('GET', url);
    xhr.onreadystatechange = function() {
        if (xhr.readyState>3 && xhr.status==200) {
            collectedData = JSON.parse(xhr.responseText);
            console.log('Data Collectd')
            getSingleCardData(collectedData, renderFunction)
        }
    };
    xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest');
    xhr.send();
    return xhr;
}


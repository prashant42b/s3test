$(document).ready(function() {
    renderContent();

    //Click event handler
    $("#myButton").click(function() {

        fetchS3Data();
        alert("Fetching AWS S3 Bucket Contents!");
       
    });
});

function renderContent()
{
    var htmlText = ``
    htmlText+=`<h1>jQuery Frontend with GoLang!</h1>
    <button id="myButton">Connect to AWS S3</button>`;

    $("#pageBody").html(htmlText);
}

function fetchS3BucketContent()
{
    
}
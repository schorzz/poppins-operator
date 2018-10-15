mermaid.initialize({
    theme: 'neutral',
    startOnLoad: false,
    sequence: {
        mirrorActors: false,
        bottomMarginAdj: 5,
        height: 50
    }
});
var specUrl = 'openapi.yaml'
Redoc.init(
    specUrl,
    {},
    document.getElementsByTagName('body')[0],
    function () {
        window.setTimeout(function (){
            mermaid.init('.diagram');
        }, 500);
    }
);
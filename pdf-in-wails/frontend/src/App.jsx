import './App.css';
import {useState} from 'react';
import {Document, Page, pdfjs} from 'react-pdf';
import {toBytes} from 'fast-base64';
import {OpenAndGetPDFData} from "../wailsjs/go/main/App";

function App() {
    pdfjs.GlobalWorkerOptions.workerSrc = `//cdnjs.cloudflare.com/ajax/libs/pdf.js/${pdfjs.version}/pdf.worker.min.js`;
    const [pdfDocument, setPdfDocument] = useState(null)
    const [numPages, setNumPages] = useState(null);
    const [pageNumber, setPageNumber] = useState(1);

    function onDocumentLoadSuccess({ numPages }) {
      setNumPages(numPages);
    }

    async function openPDF() {
        try {
            const pdfFile = await OpenAndGetPDFData()
            const pdfBytes = await toBytes(pdfFile)
            setPdfDocument({data: pdfBytes })
        } catch(e) {
            console.log("Failed to open pdf", e)
        }
    }

    return (
        <div id="App">
            <button className="btn" onClick={openPDF}>Open PDF</button>
            <Document file={pdfDocument} 
                    onLoadSuccess={onDocumentLoadSuccess} 
                    onLoadError={(error) => alert('Error while loading document! ' + error.message)}
                    onSourceError={(error) => alert('Error while retrieving document source! ' + error.message)}>
                <Page pageNumber={pageNumber} />
            </Document>
            <p>
                Page {pageNumber} of {numPages}
            </p>
        </div>
    )
}

export default App

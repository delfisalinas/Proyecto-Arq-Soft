import React, { useState } from 'react';
import axios from 'axios';
import '../assets/styles/Files.css';

const FileUpload = ({ courseId }) => {
    const [selectedFile, setSelectedFile] = useState(null);
    const [error, setError] = useState('');
    const [message, setMessage] = useState('');

    const handleFileChange = (e) => {
        setSelectedFile(e.target.files[0]);
    };

    const handleUpload = async () => {
        if (!selectedFile) {
            setError('Please select a file to upload.');
            return;
        }

        const userId = localStorage.getItem('userId');
        const token = localStorage.getItem('token');
        
        // Convert file to base64
        const reader = new FileReader();
        reader.readAsDataURL(selectedFile);
        reader.onload = async () => {
            const base64File = reader.result.split(',')[1]; // Remove the data:...base64, part

            try {
                const response = await axios.post('http://localhost:8080/files', {
                    name: selectedFile.name,
                    content: base64File,
                    userId: userId,
                    courseId: courseId
                }, {
                    headers: {
                        'Authorization': `Bearer ${token}`,
                        'Content-Type': 'application/json'
                    }
                });

                setMessage('File uploaded successfully.');
            } catch (err) {
                setError(`Error uploading file: ${err.response ? err.response.data.error : err.message}`);
            }
        };

        reader.onerror = (error) => {
            setError(`Error reading file: ${error}`);
        };
    };

    return (
        <div className="file-upload-container">
            <input type="file" onChange={handleFileChange} />
            <button onClick={handleUpload}>Upload File</button>
            {error && <p className="error-message">{error}</p>}
            {message && <p className="success-message">{message}</p>}
        </div>
    );
};

export default FileUpload;

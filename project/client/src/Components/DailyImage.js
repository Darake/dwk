import { useQuery } from 'react-query'
import axios from "axios"

const DailyImage = () => {
  const fetchImage = async () => {
    const { data } = await axios.get('/api/daily-image', { responseType: 'blob' })    
    return URL.createObjectURL(data)
  }

  const { data } = useQuery("image", fetchImage)

  return (
    <img src={data} alt="daily" width="400px" />
  )
}

export default DailyImage
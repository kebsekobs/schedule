import style from './CellStyle.css'
export const CellStyle = {
    centredCell: (content) => {
        return <div className={style['centred-cell']}>{content}</div>
    }
}
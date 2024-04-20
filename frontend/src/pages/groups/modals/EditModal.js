import React, {useEffect, useState} from 'react';
import { useForm } from 'react-hook-form';
import {useAddGroupsMutation} from "../api/AddGroupMutation";
import {useGroupByIdQuery} from "../api/getGroupById";
import {useEditGroupMutation} from "../api/EditGroupMutation";

const EditGroupModal = ({ isOpen, toggleModal, id }) => {
    const groupByIdQuery = useGroupByIdQuery(id);
    const [isDataLoaded, setIsDataLoaded] = useState(false); // Состояние для отслеживания загрузки данных

    const { register, handleSubmit, formState: { errors }, reset } = useForm({
        defaultValues: {
            groupId: '',
            name: '',
            capacity: '',
            magistracy: false
        }
    });

    const editGroupMutation = useEditGroupMutation();
    const onSubmit = data => {
        console.log(data)
        data.id = id;
        editGroupMutation.mutateAsync(data)
        toggleModal();
        reset()
    };

    useEffect(() => {
        if (groupByIdQuery.data) {
            reset({
                groupId: groupByIdQuery.data.groupId,
                name: groupByIdQuery.data.name,
                capacity: groupByIdQuery.data.capacity,
                magistracy: groupByIdQuery.data.magistracy || false
            });
            setIsDataLoaded(true); // Помечаем данные как загруженные
        }
    }, [groupByIdQuery.data, reset]);

    if (!isOpen) {
        return null;
    }

    // Рендерим загрузку, если данные еще не загружены
    if (!isDataLoaded) {
        return <div>Loading...</div>;
    }

    return (
        <div style={styles.backdrop}>
            <div style={styles.modal}>
                <h2>Введите данные</h2>
                <form onSubmit={handleSubmit(onSubmit)}>
                    <input
                        type={"number"}
                        {...register("groupId", { required: true })}
                        placeholder="Введите номер группы"
                        style={styles.input}
                    />
                    {errors.id && <span style={styles.error}>Это поле обязательно</span>}
                    <input
                        {...register("name", { required: true })}
                        placeholder="Введите название группы"
                        style={styles.input}
                    />
                    {errors.name && <span style={styles.error}>Это поле обязательно</span>}
                    <input
                        type={"number"}
                        {...register("capacity", { required: true })}
                        placeholder="Введите количество студентов"
                        style={styles.input}
                    />
                    {errors.capacity && <span style={styles.error}>Это поле обязательно</span>}
                    <label htmlFor="magistracy" style={styles.label}>
                        <input
                            type="checkbox"
                            id="magistracy"
                            {...register("magistracy")}
                            style={styles.checkbox}
                        />
                        Магистратура
                    </label>
                    <button type="submit" style={styles.button}>Отправить</button>
                </form>
                <button onClick={toggleModal} style={styles.buttonClose}>Закрыть</button>
            </div>
        </div>
    );
};


const styles = {
    backdrop: {
        position: 'fixed',
        top: 0,
        left: 0,
        right: 0,
        bottom: 0,
        backgroundColor: 'rgba(0,0,0,0.5)',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        zIndex: 1000,
    },
    modal: {
        backgroundColor: '#fff',
        padding: '20px',
        borderRadius: '8px',
        display: 'flex',
        flexDirection: 'column',
        gap: '10px',
    },
    input: {
        width: '300px',
        padding: '10px',
        margin: '10px 0',
        borderRadius: '4px',
        border: '1px solid #ccc',
    },
    checkbox: {
        width: '20px',
        height: '20px',
        margin: '10px 5px',
        cursor: 'pointer'
    },
    button: {
        padding: '10px 20px',
        border: 'none',
        backgroundColor: 'blue',
        color: 'white',
        borderRadius: '5px',
        cursor: 'pointer',
        marginTop: '10px'
    },
    buttonClose: {
        padding: '10px 20px',
        border: 'none',
        backgroundColor: 'gray',
        color: 'white',
        borderRadius: '5px',
        cursor: 'pointer',
        marginTop: '10px'
    },
    error: {
        color: 'red',
        fontSize: '14px',
        marginTop: '5px'
    },
    label: {
        display: 'flex',
        alignItems: 'center',
        cursor: 'pointer'
    }
};

export default EditGroupModal;

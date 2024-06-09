import React, { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { useAddGroupsMutation } from "../api/AddGroupMutation";
import { useGroupByIdQuery } from "../api/getGroupById";
import { useEditGroupMutation } from "../api/EditGroupMutation";
import styles from "../../shared/style/modal.module.css";
import Button from "../../../components/button";
import {CloseSvg} from "../../../components/close-svg";
import {useGroupsQuery} from "../api/getGroupsQuery";

const EditGroupModal = ({ isOpen, toggleModal, id }) => {
  const groupByIdQuery = useGroupByIdQuery(id);
  const [isDataLoaded, setIsDataLoaded] = useState(false); // Состояние для отслеживания загрузки данных
  const groupsQuery = useGroupsQuery();
  const existingGroupIds = groupsQuery.data?.map(group => group.id) ?? []
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm({
    defaultValues: {
      groupId: "",
      id: "",
      capacity: "",
      magistracy: false,
    },
  });

  const editGroupMutation = useEditGroupMutation();
  const onSubmit = (data) => {
    editGroupMutation.mutateAsync(data);
    toggleModal();
    reset();
  };

  useEffect(() => {
    if (groupByIdQuery.data) {
      reset({
        groupId: groupByIdQuery.data.groupId,
        id: groupByIdQuery.data.id,
        capacity: groupByIdQuery.data.capacity,
        magistracy: groupByIdQuery.data.magistracy || false,
      });
      setIsDataLoaded(true);
    }
  }, [groupByIdQuery.data, reset]);

  if (!isOpen) {
    return null;
  }


  if (!isDataLoaded) {
    return <div>Loading...</div>;
  }

  return (
      <div className={styles["backdrop"]}>
        <div className={styles["modal"]}>
          <div className={styles["modal-header"]}>
            <h2>Введите данные</h2>
            <Button onClick={toggleModal} styleFeature="close">
              <CloseSvg />
            </Button>
          </div>
          <form
              onSubmit={handleSubmit(onSubmit)}
              className={styles["groups-form"]}
          >
            <label>Укажите степень</label>
            <label htmlFor="consent" className={styles["label"]}>
              <input
                  type="checkbox"
                  id="magistracy"
                  className={styles["checkbox"]}
                  {...register("magistracy")}
              />
              Магистратура
            </label>
            <label>Введите номер группы</label>
            <input
                type="number"
                {...register("groupId", { required: true })}
                placeholder="118"
                className={styles["input"]}
            />
            {errors.groupId && (
                <span className={styles["error"]}>Это поле обязательно</span>
            )}
            <label>Введите код группы</label>
            <input
                {...register("id", {
                  required: "Это поле обязательно",
                  validate: value => !existingGroupIds.includes(value) || "Такой ID уже существует"
                })}
                placeholder="БOЮ15-РПИ2101"
                className={styles["input"]}
            />
            {errors.id && (
                <span className={styles["error"]}>{errors.id.message}</span>
            )}
            <label>Введите количество студентов</label>
            <input
                type="number"
                {...register("capacity", { required: true, min: 1 })}
                placeholder="13"
                className={styles["input"]}
            />
            {errors.capacity && (
                <span className={styles["error"]}>
              {errors.capacity.type === "required"
                  ? "Это поле обязательно"
                  : "Значение должно быть больше 0"}
            </span>
            )}
            <Button type="submit">Отправить</Button>
          </form>
        </div>
      </div>
  );
};

export default EditGroupModal;
